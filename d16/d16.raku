my %rules;
for lines() {
    last unless $_;
    # - negates character class, like ^
    # must quote or escape symbols and whitespace
    die unless /^ (<-[:]>+) ': ' (\d+)\-(\d+) ' or ' (\d+)\-(\d+) $/;
    %rules{$0} = ($1..$2)|($3..$4);
}

# get() parentheses needed to avoid parse error
die unless get() eq 'your ticket:';
my @mine = get.split(',');

die unless !get() && get() eq 'nearby tickets:';
# unclear when it is necessary to coerce to Int
my @nearby = lines>>.split(',')>>.Int;

# @nearby is Array which does not flatten directly
say @nearby.List.flat.grep(none(%rules.values)).sum;

# redundant with above, but select if all fields match any rule
my @valid = @nearby.grep({ all $_.map(* ~~ any(%rules.values)) });
# dense! build candidates from each field, keeping if all ticket values match
my %candidates = %rules
    .pairs
    .map({
        .key => (^@mine)
            .grep(-> $i {
                all @valid.map(*[$i]).map(* ~~ .value)
            }).SetHash
    });
my %mapping;
while %candidates {
    for %candidates.kv -> $key, $value {
        next unless $value == 1;
        # pick() randomly selects an element, but there is only one
        my $index = $value.pick;
        %mapping{$key} = $index;
        # remove resolved mapping from candidates
        .unset($index) for %candidates.values;
        %candidates{$key}:delete;
    }
}
say [*] %mapping.pairs.grep({ .key ~~ /^departure/ }).map({ @mine[.value] });
# can index into array with list
# say [*] @mine[%mapping.pairs.grep({ .key ~~ /^departure/ }).map(*.value)];
