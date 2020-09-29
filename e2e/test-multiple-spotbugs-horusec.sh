#!/bin/bash
# Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

AUTHORIZATION=$1
rm -rf ./tmp
rm -rf ./analysis
mkdir -p ./tmp
mkdir -p ./analysis

for i in $(seq 1 5)
do
  unzip -d "./analysis/spotbug-$i" ../development-kit/pkg/utils/test/zips/kotlin-spotbug/kotlin-spotbug.zip
done


for i in $(seq 1 5)
do
  horusec start -p="./analysis/spotbug-$i" -a="$AUTHORIZATION" -o="json" -O="./tmp/spotbug-$i.json" &
done
