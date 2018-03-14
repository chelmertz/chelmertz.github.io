#!/bin/bash
echo "Displaying preview of site at http://localhost:8000"
pushd _site && python -mSimpleHTTPServer; popd
