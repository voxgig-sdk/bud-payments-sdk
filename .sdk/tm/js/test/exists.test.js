
const { test, describe } = require('node:test')
const { equal } = require('node:assert')


const { BudPaymentsSDK } = require('..')


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await BudPaymentsSDK.test()
    equal(null !== testsdk, true)
  })

})
