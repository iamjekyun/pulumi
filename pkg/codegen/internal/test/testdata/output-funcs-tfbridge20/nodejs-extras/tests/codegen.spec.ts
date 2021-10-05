// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import "mocha";
import * as assert from "assert";

import * as pulumi from "@pulumi/pulumi";

import { listStorageAccountKeysOutput, ListStorageAccountKeysResult } from "../listStorageAccountKeys";

pulumi.runtime.setMocks({
    newResource: function(_: pulumi.runtime.MockResourceArgs): {id: string, state: any} {
        throw new Error("newResource not implemented");
    },
    call: function(args: pulumi.runtime.MockCallArgs) {
        if (args.token == "mypkg::listStorageAccountKeys") {
            return {
                "keys": [{
                    "creationTime": "my-creation-time",
                    "keyName": "my-key-name",
                    "permissions": "my-permissions",
                    "value": JSON.stringify(args.inputs),
                }]
            };
        }
        throw new Error("call not implemented for " + args.token);
    },
});

function checkTable(done: any, transform: (res: any) => any, table: {given: pulumi.Output<any>, expect: any}[]) {
    checkOutput(done, pulumi.all(table.map(x => x.given)), res => {
        res.map((actual, i) => {
            assert.deepStrictEqual(transform(actual), table[i].expect);
        });
    });
}

describe("output-funcs", () => {

    it("listStorageAccountKeysOutput", (done) => {
        const output = listStorageAccountKeysOutput({
            accountName: pulumi.output("my-account-name"),
            resourceGroupName: pulumi.output("my-resource-group-name"),
        });
        checkOutput(done, output, (res: ListStorageAccountKeysResult) => {
            assert.equal(res.keys.length, 1);
            const k = res.keys[0];
            assert.equal(k.creationTime, "my-creation-time");
            assert.equal(k.keyName, "my-key-name");
            assert.equal(k.permissions, "my-permissions");
            assert.deepStrictEqual(JSON.parse(k.value), {
                "accountName": "my-account-name",
                "resourceGroupName": "my-resource-group-name"
            });
        });
    });

    it("listStorageAccountKeysOutput with optional arg set", (done) => {
        const output = listStorageAccountKeysOutput({
            accountName: pulumi.output("my-account-name"),
            resourceGroupName: pulumi.output("my-resource-group-name"),
            expand: pulumi.output("my-expand"),
        });
        checkOutput(done, output, (res: ListStorageAccountKeysResult) => {
            assert.equal(res.keys.length, 1);
            const k = res.keys[0];
            assert.equal(k.creationTime, "my-creation-time");
            assert.equal(k.keyName, "my-key-name");
            assert.equal(k.permissions, "my-permissions");
            assert.deepStrictEqual(JSON.parse(k.value), {
                "accountName": "my-account-name",
                "resourceGroupName": "my-resource-group-name",
                "expand": "my-expand"
            });
        });
    });

 });


function checkOutput<T>(done: any, output: pulumi.Output<T>, check: (value: T) => void) {
    output.apply(value => {
        try {
            check(value);
            done();
        } catch (error) {
            done(error);
        }
    });
}
