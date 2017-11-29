#!/bin/bash

for i in `seq 1 2000`; do curl http://localhost:8080/hello; done