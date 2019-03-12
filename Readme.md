### Performance
You can use the `pprof`` tool to view the program's performance.

Download the latest pprof tool
  go get -u github.com/google/pprof

Run the benchmark test and create a cpu profile
  go test -bench=MonteCarlo -cpuprofile=cpuprof.out

Load the profile into pprof and open it in a web page
  pprof -http=:8080 cpuprof.out
