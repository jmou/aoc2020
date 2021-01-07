my @input = lines;

enum Dir(
    N => (0, 1),
    E => (1, 0),
    S => (0, -1),
    W => (-1, 0),
);

my @pos = 0, 0;
# Could probably simplify out the enum entirely by rotating unit vector.
# enum must be stored in $ scalar
my $dir = E;
for @input -> $line {
    $line ~~ /^(\w)(\d+)$/;
    given $0 {
        # eqv needed for structural equality; == doesn't work right
        # pred()/succ() give previous/next, but don't wrap
        when 'L' { for ^$1/90 { if $dir eqv N { $dir = W } else { $dir.=pred } } }
        when 'R' { for ^$1/90 { if $dir eqv W { $dir = N } else { $dir.=succ } } }
        # X* cross-product meta operator with scalar multiplies each element
        # >>*>> multiplication hyperoperator doesn't work here
        # >>+=>> hyperoperator applies to each element
        # points at smaller list, but with equal lengths >>+=<< also valid
        # (<<+=>> is "do what I mean" (dwimmy) and automatically lengthens)
        when 'F' { @pos >>+=>> (@($dir) X* $1) }
        # enums() gives map of variants
        # >>*>> works here instead of X*
        default { @pos >>+=>> (Dir.enums{$0} X* $1) }
    }
}
# Manhattan distance
say @pos>>.abs.sum;

@pos = 0, 0;
my @waypoint = 10, 1;
for @input -> $line {
    $line ~~ /^(\w)(\d+)$/;
    given $0 {
        # given as a statement modifier topicalizes
        when 'L' { for ^$1/90 { @waypoint = -.[1], .[0] given @waypoint } }
        when 'R' { for ^$1/90 { @waypoint = .[1], -.[0] given @waypoint } }
        when 'F' { @pos >>+=>> (@waypoint X* $1) }
        default { @waypoint >>+=>> (Dir.enums{$0} X* $1) }
    }
}
say @pos>>.abs.sum;
