
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { BudPaymentsSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = BudPaymentsSDK.test()
    equal(testsdk instanceof BudPaymentsSDK, true,
      'BudPaymentsSDK.test() must return a client synchronously')
  })

})
