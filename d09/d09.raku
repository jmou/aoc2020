# Int necessary for minmax() to work correctly
my @input = lines>>.Int;

my @window = @input[^25];
my $invalid;
for @input[25..*] -> $num {
    $invalid = $num and last unless @window.combinations(2).first(*.sum == $num);
    @window.shift;
    @window.push($num);
}
say $invalid;

# X= cross-product metaoperator with assignment initialization trick
my ($min, $max) X= 0;
loop {
    # <=> comparison returns Order enum
    given @input[$min..$max].sum <=> $invalid {
        when Same { last }
        when More { $min++ }
        when Less { $max++ }
    }
}
# int-bounds() to sum min + max instead of every integer in range
say @input[$min..$max].minmax.int-bounds.sum;