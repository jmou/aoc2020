my @input = lines();
# + coerces array to int; gives count, same as .elems()
# grep smartmatch filters
# topic variable $_ is default parameter for block w/o signature
# :s makes regex whitespace sensitive
# comb returns matches
# ~~ smartmatch calls RHS.ACCEPTS(LHS)
say +@input.grep({m:s/^(\d+)\-(\d+) (\w)\: (\w+)$/ and +$3.comb($2) ~~ ($0..$1)});
# ~ coerces to a string
# TODO why do we need it? (otherwise it seems to always accept)
say +@input.grep({m:s/^(\d+)\-(\d+) (\w)\: (\w+)$/ and +$3.comb[$0-1,$1-1].grep(~$2) == 1});
# TODO how to DRY?
