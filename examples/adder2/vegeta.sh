echo "GET http://localhost:8080/add/1/2" | vegeta attack -duration=10s | tee results.bin | vegeta report
