import { test, expect } from "@playwright/test";
import { getEndpointDefinition } from '../utils';
import { generateFastenConnectAuthorizeUrl, generateSourceAuthorizeUrl } from '@shared-library';

test("Qualifacts-Credible Login Flow", async ({ page }, testInfo) => {
    try {
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: "setSessionName", arguments: { name: testInfo.title } })}`);
        await page.waitForTimeout(5000);

        //get the Cerner Sandbox endpoint definition
        // let endpointDefinition = await getEndpointDefinition('ac8308d1-90de-4994-bb3d-fe404832714c')
        // let authorizeData = await generateSourceAuthorizeUrl(endpointDefinition)
        let authorizeData = await generateFastenConnectAuthorizeUrl(
            'ba036673-b531-4e3b-a099-3cf5c024a71d',
            '646a25b6-1d77-406a-b8ee-b6093556800d',
            '0f2937c9-4cfc-4594-8f5a-8c878888c827',
        )


        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());
        await page.getByRole('heading', { name: 'Log in to your member account' }).waitFor({ state: 'visible' });
        await page.click('input.MuiOutlinedInput-input');
        await page.keyboard.type("HOSPatient");
        await page.click('input.MuiInputBase-inputAdornedEnd');
        await page.keyboard.type("HOSPatient2023");
        await page.click("button[type='submit']");


        await page.waitForSelector('div:has-text("Whose record do you want to allow access to?")', { state: 'visible' });
        await page.click('button:has-text("Linda558 Hessel84 (31 years)")');
        await page.click('button:has-text("Agree")');
        await page.waitForSelector("text=Example Domain");





        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'passed', reason: 'Authentication Successful' } })}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'failed', reason: 'Test failed' } })}`);
    }
});
