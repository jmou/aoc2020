my @input = lines>>.Int.sort;

my %diff;
for @input.rotor(2 => -1) -> ($lo, $hi) {
    %diff{$hi - $lo}++;
}
say [*] %diff{1, 3}.map(* + 1);

# Bag is a set that associates to counts
my $bag = @input.rotor(2 => -1).map({ .[1] - .[0] }).Bag;
say [*] $bag{1, 3}.map(* + 1);

my @arrangements is default(0) = 1;
for ^@input -> $i {
    my $jolts = @input[$i];
    @arrangements[$jolts] = @arrangements[$jolts-1];
    @arrangements[$jolts] += @arrangements[$jolts-2] if $jolts >= 2;
    @arrangements[$jolts] += @arrangements[$jolts-3] if $jolts >= 3;
}
say @arrangements.tail;