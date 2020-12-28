my @required-fields = <byr iyr eyr hgt hcl ecl pid>;

# lines(:nl-in("\n\n")) doesn't seem to work on $*ARGFILES or $*IN, presumably
# because they are already IO::Handle (or IO::CatHandle), not IO::Path.
# Setting .nl-in works though
# $*ARGFILES.nl-in = "\n\n";
# say $*ARGFILES.lines;

# gather-take returns a sequence, like a generator
# say [+] gather for $*ARGFILES.slurp.split("\n\n") {
#     # .trim for trailing \n on last paragraph
#     my @fields = (.split(':')[0] for .trim.split(/\s+/));
#     take 1 if @required-fields (<=) @fields;
# }

# $*ARGFILES reads from arguments if given otherwise stdin; like perl <>
# .trim for trailing \n on last paragraph
# % hash coercion on even number of elements
# split records by paragraph, fields by whitespace, key/value by :
my @passports = (%(|.split(':') for .trim.split(/\s+/)) for $*ARGFILES.slurp.split("\n\n"));

# (>=) superset operator
my @p1 = @passports.grep: { .keys (>=) @required-fields };
say @p1.elems;

# if we initialize with { } hash composer, raku complains it is unnecessary
# (it suggests using the := binding operator)
# is default will be returned for nonexistent keys
# True.ACCEPTS anything (i.e., any unknown field is valid)
# (without is default, the implicit constraint Any type object would be returned
# instead, which also ACCEPTS anything)
my %validators is default(True) =
    # <= etc can be chained
    byr => 1920 <= * <= 2002,
    iyr => 2010 <= * <= 2020,
    eyr => 2020 <= * <= 2030,
    # ?? !! is ternary operator
    # ** general quantifier, like {2,6} repetition
    hgt => { m/^(\d**2..3)(cm|in)$/ && ($1 eq 'cm' ?? 150 <= $0 <= 193 !! 59 <= $0 <= 76) },
    # \# must be escaped (else is a comment)
    # https://github.com/rakudo/rakudo/issues/4142 escaping bug
    # <[]> character class
    hcl => /^ '#' <[\da..f]> ** 6 $/,
    # any Junction autothreads over eigenstates; collapses in Boolean context
    # <> word quoting returns list of strings (actually allomorphs)
    ecl => any(<amb blu brn gry grn hzl oth>),
    pid => /^ \d ** 9 $/;

say +@p1.grep: -> %fields {
    # Junction form (not short circuiting)
    # all((-> $kv { $kv.value ~~ %validators{$kv.key} } for %fields));

    # first() gives short-circuiting behavior
    # $_ topic variable doesn't work on RHS of ~~ because it binds $_ to LHS
    # ! negates any relational operator
    # {} postcircumfix subscripting
    !%fields.first: -> $kv { $kv.value !~~ %validators{$kv.key} };
};

# grep() over passports
# map() each to generate a lazy Seq, which validates each field
# [] reduction metaoperator applies infix to a list
# [&&] checks all valid and should short circuit
say +@p1.grep: { [&&] .map(-> $kv { $kv.value ~~ %validators{$kv.key} }) };
