# first time looking at performance. use --profile=prof.html
# can't interrupt with ctrl-c
# millions of allocations; probably main source of slowness
# profiles actually quite hard to interpret, janky to navigate
# hard to constrain types of nested arrays; not sure how to optimize

# unit-scoped MAIN applies to rest of file
# sub MAIN bootstraps a CLI with flags generated from parameters; can be multi
# | any operator returns Junction
# Int MAIN parameter silently accepts cast from Bool without number, like --part
unit sub MAIN(Int :$part where 1|2 = 1, Bool :$print);

# $*IN is stdin (bare lines() inside sub MAIN crashes; bug?)
# $* dynamic variable twigil looks up through calling scopes (not lexical)
my @grid = $*IN.lines.map: { flat 'x', .comb, 'x' };
# xx list repetition to fill borders to avoid bounds checking
@grid.unshift('x' xx @grid[0]);
@grid.push('x' xx @grid[0]);

# int native type 20% speedup (avoids containers on arithmetic, like boxing)
# multi dispatch adds ~50% overhead
# 23 sec
multi sub compute-cell(int $r, int $c, 1) {
    # [$r;$c] indexing syntax is valid but slower
    # := binding saves ~25%
    my $old := @grid[$r][$c];
    return $old if $old eq '.'|'x';

    # refactoring out this constant saves ~15%
    # constant deltas = [(-1..1 X -1..1).grep: * !eqv (0, 0)];
    # my $neighbors = (@grid[$r + .[0]][$c + .[1]] eq '#' for deltas).sum;
    # hardcoding this logic ~2x faster! presumably by avoiding allocations
    my int $neighbors = 0;
    $neighbors++ if @grid[$r-1][$c-1] eq '#';
    $neighbors++ if @grid[$r-1][$c] eq '#';
    $neighbors++ if @grid[$r-1][$c+1] eq '#';
    $neighbors++ if @grid[$r][$c-1] eq '#';
    $neighbors++ if @grid[$r][$c+1] eq '#';
    $neighbors++ if @grid[$r+1][$c-1] eq '#';
    $neighbors++ if @grid[$r+1][$c] eq '#';
    $neighbors++ if @grid[$r+1][$c+1] eq '#';

    return '#' if $old eq 'L' && $neighbors == 0;
    return 'L' if $old eq '#' && $neighbors >= 4;
    return $old;
}

# 1 min 47 sec
multi sub compute-cell(int $r, int $c, 2) {
    my $old := @grid[$r][$c];
    return $old if $old eq '.'|'x';

    constant deltas = [(-1..1 X -1..1).grep: * !eqv (0, 0)];
    my int $neighbors = 0;
    for deltas -> (int $dr, int $dc) {
        my (int $r0, int $c0) = $r, $c;
        # repeat is like do-while
        repeat {
            $r0 += $dr;
            $c0 += $dc;
        } while @grid[$r0][$c0] eq '.';
        $neighbors++ if @grid[$r0][$c0] eq '#';
    }

    return '#' if $old eq 'L' && $neighbors == 0;
    return 'L' if $old eq '#' && $neighbors >= 5;
    return $old;
}

loop {
    # cache() avoids issues with consumed Seq
    # Array seems to need it too sometimes though
    (say .cache.join for @grid) if $print;
    # pre-allocating and shaped array don't seem to improve performance
    # shaped array doesn't seem well supported (no eqv structural equality?)
    my @next = [[] xx @grid];  # allocate just enough to avoid racy writes
    # int native type shaves a few %
    # race parallelism ~30% speedup; method form needed to specify tuning
    (^@grid).race(:8batch).map: -> int $r {
        for ^@grid[$r] -> int $c {
            @next[$r][$c] = compute-cell($r, $c, $part);
        }
    }
    # eqv equivalence operator checks structure recursively
    last if @grid eqv @next;
    @grid = @next;
}

# flat() doesn't flatten Arrays (need to decont elements); slip() instead
say @grid.map(&slip).grep('#').elems;
