my @input = lines();

# topic variable $_ is default parameter for block w/o signature
# say +@input.grep({m:s/^(\d+)\-(\d+) (\w)\: (\w+)$/ and +$3.comb($2) ~~ ($0..$1)});

# grep smartmatch filters
# :s makes regex whitespace sensitive
# Making this a sub doesn't seem to work; $0, etc capture group variables in
# &filter are Nil.
my &filter-input = -> &filter { @input.grep({m:s/^(\d+)\-(\d+) (\w)\: (\w+)$/ and &filter()}) };

# + coerces array to int; gives count, same as .elems()
# regex captures are 0-indexed
# comb returns matches
# ~~ smartmatch calls RHS.ACCEPTS(LHS)
say +&filter-input({ +$3.comb($2) ~~ ($0..$1) });
# ~ coerces to a string (Match.ACCEPTS returns self)
# Returning self kind of makes sense to allow $foo ~~ m/bar/
say +&filter-input({ +$3.comb[$0-1,$1-1].grep(~$2) == 1 });
