# grammar are essentially a collection of named regexes
grammar Rule {
    # rule use :s (:sigspace) which uses <.ws> for unquoted whitespace
    rule TOP { <type> bags contain <contents> '.' }

    # proto is a formalization over multi methods, used here for alternation
    proto rule contents {*}
    # :sym<> adverb for each alternation
    rule contents:sym<nothing> { 'no other bags' }
    # % modifies the + quantifier to intersperse a delimiter
    rule contents:sym<something> { ( <num> <type> bags? )+ % ',' }

    # token do not use :s
    token num { \d+ }
    token type { \w+ ' ' \w+ }
}

# actions are just normal classes used as visitors while parsing a grammer
class RuleActions {
    # $/ as the parameter is a convention to use $<> match variables
    # make() stores data on an AST node to be retrieved by made()
    method TOP($/) { make $<type> => $<contents>.made }
    # hash() to convert list of pairs to Associative
    # also could have used the obviously apt Bag
    method contents:sym<something>($/) { make $0.map({.<type> => .<num>}).hash }
}

# \ sigilless variable does not have a container and aliases a value permanently
sub parse(\lines) {
    lines.map({ Rule.parse($_, actions => RuleActions).made }).hash
}

sub count-contains-gold(%bags) {
    # constraint value type object is default for nonexistent
    # undefined values are not null; they are their type object
    my Array %reverse;
    for %bags.kv -> $outer, $contents {
        for $contents.kv -> $inner, $num {
            # push() assumes default Array type object to autovivify
            %reverse{$inner}.push($outer)
        }
    }

    # mutating method call on type object to instantiate itself
    # SetHash is mutable Set variant
    my SetHash $visited .= new;
    sub visit($type) {
        # (cont) equivalent to ∋
        return if $visited (cont) $type;
        $visited.set($type);
        # .& invokes subroutine with invocant as first parameter
        # :v adverb skips nonexistent values (will return empty List)
        # also will decont from Scalar container so value behaves as a list
        # otherwise would need <> Zen slice to explicitly decont
        .&visit for %reverse{$type}:v;
    }
    visit('shiny gold');
    return $visited.elems - 1;  # exclude shiny gold
}

sub count-bags(%bags, $type) {
    1 + (count-bags(%bags, .key) * .value for %bags{$type}.pairs).sum
}

my %bags = parse(lines);
say count-contains-gold(%bags);
say count-bags(%bags, 'shiny gold') - 1;  # exclude shiny gold

# DOC CHECK phaser runs on `raku --doc -c`
DOC CHECK {
    # raku outputs TAP results
    use Test;
    subtest {
        # :to heredoc customizes quoting terminator; indentation is removed
        my $in = q:to/END/;
        light red bags contain 1 bright white bag, 2 muted yellow bags.
        dark orange bags contain 3 bright white bags, 4 muted yellow bags.
        bright white bags contain 1 shiny gold bag.
        muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
        shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
        dark olive bags contain 3 faded blue bags, 4 dotted black bags.
        vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
        faded blue bags contain no other bags.
        dotted black bags contain no other bags.
        END
        # is tests with ==
        is count-contains-gold(parse($in.lines)), 4;
    }
    subtest {
        my $in = q:to/END/;
        shiny gold bags contain 2 dark red bags.
        dark red bags contain 2 dark orange bags.
        dark orange bags contain 2 dark yellow bags.
        dark yellow bags contain 2 dark green bags.
        dark green bags contain 2 dark blue bags.
        dark blue bags contain 2 dark violet bags.
        dark violet bags contain no other bags.
        END
        is count-bags(parse($in.lines), 'shiny gold') - 1, 126;
    }
    # finish TAP testing; plan with number of tests is preferred
    done-testing;
}
