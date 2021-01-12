# get() reads one line
my $start = get;
my @schedule = get.split: ',';

# div apparently not defined on Cool, need to coerce to Int
say gather for @schedule.grep(* ne 'x')>>.Int -> $bus {
    my $departure = (($start - 1) div $bus + 1) * $bus;
    take { :$departure, val => $bus * ($departure - $start) };
}.min(*<departure>)<val>;

my ($rem, $div) = 0, 1;
# pairs() less ergonomic to iterate, but kv() destructures elements
for @schedule.pairs.grep({ .value ne 'x' })>>.kv -> ($i, $bus) {
    # %% divisibility operator means $a % $b == 0
    $rem += $div until ($rem + $i) %% $bus;
    # I think this is where LCM is supposed to be used
    $div lcm= $bus;
};
say $rem;
