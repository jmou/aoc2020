my @initial = @*ARGS.join.split(',');
# using array about twice as fast as hash
# native int about 5x faster (but can't use :exists adverb)
my int @said;
# $ anonymous state variable initialized once, like static
@said[$_] = ++$ for @initial[0..*-2];
my $num = @initial[*-1];

# takes several minutes to run
for (@initial..^2020, 2020..^30000000) {
    for $_ {
        my $next = @said[$num] ?? $_ - @said[$num] !! 0;
        @said[$num] = $_;
        $num = $next;
    }
    say $num;
}
