#!/bin/bash
for i in {1..50}
do
   ./kpmp -i Instances/set1/cc2.txt -o Solutions/cc2/solution%%n.txt -k 3 -n 50000 -s -c stats.csv
done