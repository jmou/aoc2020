# .list makes the Seq into a lazy List
# := binding keeps @input as a List
# = can be used and would create an Array from the Seq or List
# my @input := 'd01/input'.IO.lines.list;
# prefix + hyperoperator treats as Int
my Int @input = +<< 'd01/input'.IO.lines;
# reduction metaoperator with *
# {say [*] $_; last} if .sum == 2020 for @input.combinations(2);
# * Whatever-currying creates a Callable closure
say [*] @input.combinations(2).first(*.sum == 2020);
say [*] @input.combinations(3).first(*.sum == 2020);
