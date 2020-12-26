my @groups = slurp.split("\n\n");

# sum() evaluates in scalar context, so gathered lists become sizes
say gather for @groups { take .comb(/\w/).unique }.sum;

say gather for @groups -> $group {
    # (&) set intersection used to [] reduce gathered sequence of answers
    take [(&)] gather for $group.trim.split("\n") -> $answers {
        take $answers.comb(/\w/)
    }
}.sum;

# or if we're golfing
# >> postfix unary hyper operator applies operator to each element of list, like map
# (points away from the list)
say ([(&)] .trim.split("\n")>>.comb(/\w/) for @groups).sum;
