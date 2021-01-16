sub neighbors($p) {
    my @dims = $p.split(',');
    my @neighbors = [X] @dims.map({ $_ - 1, $_, $_ + 1 });
    @neighbors>>.join(',').grep(* ne $p)
}

my @input = lines;

# represent 3D with one extra dimension; 4D with two
for '0', '0,0' -> $z {
    my $active = gather for @input.kv -> $x, $line {
        for $line.comb.kv -> $y, $cell {
            take "$x,$y,$z" if $cell eq '#';
        }
    }.Set;

    for ^6 {
        # iterate the set of neighbors of active points
        # race about ~2x speedup
        my $next = race for $active.keys.map({ |neighbors($_) }).Set.keys -> $p {
            my $active-neighbors = neighbors($p).grep(* (elem) $active).elems;
            if $active{$p} {
                $p if $active-neighbors == 2|3;
            } else {
                $p if $active-neighbors == 3;
            }
        }.Set;
        $active = $next;
    }

    say +$active;
}
