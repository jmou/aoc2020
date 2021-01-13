my $mask;
my (%mem-p1, %mem-p2);
for lines() -> $line {
    my ($lhs, $rhs) = $line.split(' = ');
    if $lhs eq 'mask' {
        $mask = $rhs;
    } else {
        my $value = sprintf '%036b', $rhs;
        # zip $mask and $value and backfill 'X'
        %mem-p1{$lhs} = ($mask.comb Z $value.comb).map(*.first: none 'X').join;

        $lhs ~~ /^mem\[(\d+)\]$/;
        my $addr = sprintf '%036b', $0;
        # funky ==> feed to get addresses as list of options for each digit
        $mask.comb Z $addr.comb
        ==> map(-> ($m, $v) { given $m {
            when 0 { [$v] }
            when 1 { [1] }
            when 'X' { [0, 1] }
        } })
        ==> my @addr;
        # [X] cross product reduce list of options for each digit
        # gives list of digits (as list of chars)
        %mem-p2{$_.join} = $rhs for [X] @addr;
    }
}
say %mem-p1.values.map({ :2($_) }).sum;
say %mem-p2.values.sum;

# Basically same as [X] cross product reduction
# sub enumerate(@alternatives) {
#     return '' unless @alternatives;
#     my ($head, @rest) = @alternatives;
#     return (|enumerate(@rest).map($_ ~ *) for $head<>);
# }
