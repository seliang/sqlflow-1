// Copyright 2019 The SQLFlow Authors. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testdata

// FeatureDericationCaseSQL is .sql format data samples to test feature derivation.
var FeatureDericationCaseSQL = `CREATE DATABASE IF NOT EXISTS feature_derivation_case;
DROP TABLE IF EXISTS feature_derivation_case.train;
CREATE TABLE feature_derivation_case.train (
       c1 float,
       c2 float,
       c3 TEXT,
       c4 TEXT,
       c5 TEXT,
       c6 TEXT,
       class int);
INSERT INTO feature_derivation_case.train VALUES
(6.4,2.8, '1,4,2,3', '1,3,2,6', '3,140', 'MALE', 0),
(5.0,2.3, '1,3,8,3', '3,2,5,3', '93,12,1,392,49,13,398', 'FEMALE', 1),
(4.9,2.5, '9,2,2,2', '1.2,4.8,3.2,1', '10,11,32,32,1', 'FEMALE', 1),
(5.1,2.2, '2,1,8,5', '5.0,3,2,1', '23,22,1', 'FEMALE', 1),
(4.8,3.1, '3,3,2,6', '3,2,3,5', '30,3,1,32', 'NULL', 0);`
