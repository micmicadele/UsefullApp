use strict;
use warnings;
use Prima qw(Application Buttons Edit);

my $main = Prima::MainWindow->new(
    text    => 'Calculator',      
    size    => [470, 380],        
    centered => 1,                 
    borderStyle => bs::Dialog,     
);

my $num1 = $main->insert(Edit =>
    pack => { side => 'top', fill => 'x', padx => 10, pady => 5 },
    hint => 'Enter first number',
    text => '',
);

my $num2 = $main->insert(Edit =>
    pack => { side => 'top', fill => 'x', padx => 10, pady => 5 },
    hint => 'Enter second number',
    text => '',
);

my $result = $main->insert(Edit =>
    pack => { side => 'top', fill => 'x', padx => 10, pady => 5 },
    text => 'Result',
    readonly => 1,
);

$main->insert(Button =>
    text => 'Add',
    pack => { side => 'left', fill => 'x', padx => 10, pady => 5 },
    onClick => sub {
        my $sum = $num1->text + $num2->text;
        $result->text($sum);
    },
);

$main->insert(Button =>
    text => 'Subtract',
    pack => { side => 'left', fill => 'x', padx => 10, pady => 5 },
    onClick => sub {
        my $difference = $num1->text - $num2->text;
        $result->text($difference);
    },
);

$main->insert(Button =>
    text => 'Multiply',
    pack => { side => 'left', fill => 'x', padx => 10, pady => 5 },
    onClick => sub {
        my $product = $num1->text * $num2->text;
        $result->text($product);
    },
);

$main->insert(Button =>
    text => 'Divide',
    pack => { side => 'left', fill => 'x', padx => 10, pady => 5 },
    onClick => sub {
        my $dividend = $num2->text;
        if ($dividend == 0) {
            $result->text('Error: Division by zero!');
        } else {
            my $quotient = $num1->text / $dividend;
            $result->text($quotient);
        }
    },
);

run Prima;