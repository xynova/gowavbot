#!/bin/bash
git clone https://github.com/ggerganov/ggwave --recursive
cd ggwave && mkdir -p build && cd build && cmake .. -DGGWAVE_BUILD_EXAMPLES=OFF -DCMAKE_BUILD_TYPE=Release && make
