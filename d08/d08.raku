class Machine {
    # has declares attributes (instance variables); all are private
    # $. twigil generates public accessor getter (is rw also creates setter)
    # new() default constructor takes named parameters for $. attributes
    # is required trait throws if not provided
    has @.program is required;
    # initially undefined but will be autovivified
    has SetHash $!visited;
    # $! is a "normal" attribute without any accessors (totally private)
    # defaults are evaluated at instantiation
    has $!pc = 0;
    has $.acc = 0;

    # BUILD overrides object construction to initialize attributes
    # : named parameter with arguments to new()
    # submethod is not inherited
    submethod BUILD(:@program) {
        # words() splits on whitespace
        # make everything Array, so they are containerized (mutable) for repair
        @!program = @program>>.words>>.Array;
    }

    # method here behaves like properties
    # returns (or of) define return type constraints
    # {} postcircumfix autovivifies SetHash (unlike (cont))
    method visited returns Bool { $!visited{$!pc} }
    # --> return type constraint syntax; can take :D :U
    # TODO :D can still return Nil?
    method terminated(--> Bool) { $!pc == @!program }

    # ! private method
    method !reset {
        # $! twigil for use in method, even if attribute declared with $.
        $!pc = $!acc = 0;
        $!visited = Nil;
    }

    method step {
        # {} postcircumfix assignment like set() but autovivifies
        $!visited{$!pc} = True;
        my ($op, $arg) = @.program[$!pc];

        # # given/when topicalize and smartmatch, like switch/case
        # given $op {
        #     when 'acc' { $!acc += $arg }
        #     when 'jmp' { $!pc += $arg; return }  # skip $!pc++
        #     when 'nop' {}
        #     # !!! (or ...) stub code throws when executed
        #     # (die perhaps more relevant; stubs intended for OOP)
        #     default { !!! }
        # }
        # $!pc++;

        self.compute($op, $arg);
    }

    # multi-dispatch selects the appropriate function by arguments
    multi method compute('acc', $arg) { $!acc += $arg; $!pc++; }
    multi method compute('jmp', $arg) { $!pc += $arg; }
    multi method compute('nop', $arg) { $!pc++; }

    method repair {
        # constant \ for sigilless is optional; = is actually :=
        # (sigil can be used but won't have a container)
        # {} hash constructor
        constant swap-op = { jmp => 'nop', nop => 'jmp' };
        # is rw parameter trait binds container, like passing by reference
        for @!program -> ($op is rw, $arg) {
            next unless swap-op (cont) $op;
            # self is the method invocant
            # ! instead of . for private method call
            self!reset;
            $op = swap-op{$op};
            self.step until self.visited || self.terminated;
            return if self.terminated;
            $op = swap-op{$op};
        }
    }
}

# named arguments specified by Pair values
my $machine = Machine.new: program => lines;
$machine.step until $machine.visited;
say $machine.acc;

$machine.repair;
say $machine.acc;

DOC CHECK {
    use Test;
    my @program = q:to/END/.lines;
    nop +0
    acc +1
    jmp +4
    acc +3
    jmp -3
    acc -99
    acc +1
    jmp -4
    acc +6
    END
    # : variable prefix uses same named variable as named argument
    my $machine = Machine.new: :@program;
    $machine.step until $machine.visited;
    is $machine.acc, 5;

    $machine.repair;
    is $machine.acc, 8;
}
