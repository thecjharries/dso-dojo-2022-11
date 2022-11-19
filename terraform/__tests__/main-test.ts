// Copyright (c) HashiCorp, Inc
// SPDX-License-Identifier: MPL-2.0
import "cdktf/lib/testing/adapters/jest"; // Load types for expect matchers
import { Testing } from "cdktf";
import { DsoDojo202211 } from "../main";

describe("My CDKTF Application", () => {
    describe("Checking validity", () => {
        it("check if the produced terraform configuration is valid", () => {
            const app = Testing.app();
            const stack = new DsoDojo202211(app, "test");
            expect(Testing.fullSynth(stack)).toBeValidTerraform();
        });

        it("check if this can be planned", () => {
            const app = Testing.app();
            const stack = new DsoDojo202211(app, "test");
            expect(Testing.fullSynth(stack)).toPlanSuccessfully();
        });
    });
});
