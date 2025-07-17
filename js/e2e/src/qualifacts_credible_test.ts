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
            'e5079d5c-4526-4b03-a5d9-55db63065f94',
            'a89fc38e-a602-4ef3-9e9d-2f2e1ee218e2',
            '60a12b43-00d2-4830-89df-7a063a29e39c',
        )


        // authorizeData.sourceState
        console.log(authorizeData.url.toString())

        // Start login flow by clicking on button with text "Login to MyChart"
        await page.goto(authorizeData.url.toString());

        await page.waitForSelector('#Username', { state: 'visible' });
        await page.click('#Username');
        await page.keyboard.type("jason@fastenhealth.com");
        await page.click('#Password');
        await page.keyboard.type("zs$JwSScA5o$6Xnp");
        await page.waitForSelector('button[name="button"][value="login"]', { state: 'visible' });
        await page.click('button[name="button"][value="login"]');

        try {
            const gotItButton = await page.waitForSelector('a.cc-btn.cc-dismiss', { timeout: 3000 });
            await gotItButton.click();
        } catch (e) {
            // Popup not found — ignore
        }



        await page.waitForSelector('button[name="button"][value="yes"]', { state: 'visible' });
        await page.click('button[name="button"][value="yes"]');


        await page.waitForSelector("text=Your account has been securely connected to FASTEN.")

        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'passed', reason: 'Authentication Successful' } })}`);
    } catch (e) {
        console.log(e);
        await page.evaluate(_ => { }, `browserstack_executor: ${JSON.stringify({ action: 'setSessionStatus', arguments: { status: 'failed', reason: 'Test failed' } })}`);
    }
});
