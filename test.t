use Test::More tests=>2;

system "rm -f output/*";
system " go run main.go < input/GoogleKeepDocument-tidy.html";
$count = `ls output/*|wc -l`;
$count =~ s/^\s+|\s+$//g;
is( $count , 659, "correct # of files");
is((system "grep -i prabha output/0000.html"), 0, "test output contains prabha");