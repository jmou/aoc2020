my @seat-ids = lines.map(-> $line {
    # TODO making $_ mutable
    $_ = $line;
    s:g/F|L/0/;
    s:g/B|R/1/;
    # radix form Int parse
    :2($_);
}).sort;
# or since already sorted, @seat-ids[*-1]
say @seat-ids.max;

# find discontinuity (seat id increases by more than 1) and decrement
say @seat-ids.pairs.first({ .value != .key + @seat-ids[0] }).value - 1;
