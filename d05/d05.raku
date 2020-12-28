# is copy parameter trait makes a writable copy
# my @seat-ids = lines.map(-> $_ is copy {
#     s:g/F|L/0/;
#     s:g/B|R/1/;
#     # radix form Int parse
#     :2($_);
# }).sort;
# trans substitutes characters, like tr
# parse-base() more general than Int parse
# my @seat-ids = lines.map({.trans(<F L> => '0', <B R> => '1').parse-base(2)}).sort;
# TR/// non-destructive transliteration
my @seat-ids = lines.map({ :2(TR/FLBR/0011/) }).sort;
# or since already sorted, @seat-ids[*-1]
say @seat-ids.max;

# pairs() on list gives index => value
# find discontinuity (seat id increases by more than 1) and decrement
# say @seat-ids.pairs.first({ .value != .key + @seat-ids[0] }).value - 1;

# rotor() returns sublists of specified size
# in => Pair form, key is the size and value is the gap; negative is overlap
# .[] postcircumfix implicit on $_ topic variable
say @seat-ids.rotor(2 => -1).first({.[0] + 1 != .[1]})[0] + 1;

# minmax() creates a bounding range
# (elem) equivalent to ∈
say @seat-ids.minmax.first(* !(elem) @seat-ids);

# (-) set difference
# Set does Associative role, so keys() gives elements; could also pick()
say (@seat-ids.minmax (-) @seat-ids).keys[0];
