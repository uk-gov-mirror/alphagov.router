const { test } = require('./fixtures');
const { expect } = require('@playwright/test');

test.describe('460 Gone Tests', () => {
  test('460 Gone page returns correct status code', async ({ page }) => {
    const response = await page.goto('/government/topics/transport');

    // Verify 460 status code is returned
    expect(response.status()).toBe(460);

    // Note: GOV.UK currently returns an empty response body for 460 pages
    // These are detected by the nginx layer and rewritten with the correct
    // status code (410), and a static error page from an S3 bucket.
  });
});
