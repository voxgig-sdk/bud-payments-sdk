
const envlocal = __dirname + '/../../../.env.local'
require('../../utility').loadEnvLocal(envlocal)

const Path = require('node:path')
const Fs = require('node:fs')

const { test, describe, afterEach } = require('node:test')
const assert = require('node:assert')
const { createLiveTransport } = require('../../live-runner')
const { runLiveEntity } = require('../../live-entity')


const { BudPaymentsSDK, BaseFeature, stdutil, config } = require('../../..')

const {
  envOverride,
  liveClientOptions,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
} = require('../../utility')


describe('InitiatePaymentBudLicenseEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when BUD_PAYMENTS_TEST_LIVE=TRUE.
  afterEach(liveDelay('BUD_PAYMENTS_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = BudPaymentsSDK.test()
    const ent = testsdk.InitiatePaymentBudLicense()
    assert(null != ent)
  })


  test('basic', async (t) => {

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"bud_pay_url":{"a":true,"h":"Bud Pay Url","n":"bud_pay_url","r":true,"sh":"A Bud Pay URL, allowing Customers to complete (authenticate) a relevant payment resource with the relevant payment service provider.","t":"`$STRING`","key$":"bud_pay_url","index$":0},"country_code":{"a":true,"h":"Country Code","n":"country_code","r":true,"sh":"Provider country code following the ISO 3166 alpha-3 code format","t":"`$STRING`","key$":"country_code","index$":1},"display_name":{"a":true,"h":"Display Name","n":"display_name","r":true,"sh":"The name of the payment provider to be used for display purposes (e.g.","t":"`$STRING`","key$":"display_name","index$":2},"icon":{"a":true,"h":"Icon","n":"icon","r":true,"sh":"A URL for the icon of the corresponding provider","t":"`$STRING`","key$":"icon","index$":3},"implementation_type":{"a":true,"h":"Implementation Type","n":"implementation_type","r":true,"sh":"The implementation standard adopted by the provider","t":"`$STRING`","key$":"implementation_type","index$":4},"maintenance_status":{"a":true,"h":"Maintenance Status","n":"maintenance_status","r":true,"sh":"The status of the maintenance window associated with the provider","t":"`$STRING`","key$":"maintenance_status","index$":5},"maintenance_window":{"a":true,"h":"Maintenance Window","n":"maintenance_window","r":true,"t":"`$ANY`","union":{"branches":2,"count":1,"depth":0},"key$":"maintenance_window","index$":6},"payment_details":{"a":true,"h":"Payment Details","n":"payment_details","r":true,"sh":"Payment details required to make a Single Payment","t":"`$OBJECT`","key$":"payment_details","index$":7},"payment_id":{"a":true,"h":"Payment Id","n":"payment_id","r":true,"sh":"Unique identifier for the payment","t":"`$STRING`","key$":"payment_id","index$":8},"provider":{"a":true,"h":"Provider","n":"provider","r":true,"sh":"The name (identifier) of the payment provider","t":"`$STRING`","key$":"provider","index$":9},"provider_types":{"a":true,"h":"Provider Types","n":"provider_types","r":false,"sh":"One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow.","t":"`$ARRAY`","key$":"provider_types","index$":10},"redirect_url":{"a":true,"h":"Redirect Url","n":"redirect_url","r":true,"sh":"URL where the user will be redirected to once they have completed the TPP Connection flow.","t":"`$STRING`","key$":"redirect_url","index$":11},"required_actions":{"a":true,"h":"Required Actions","n":"required_actions","r":true,"sh":"An array of enum fields to indicate if any additional actions are required in order to complete a given payment flow for the given provider.","t":"`$ANY`","key$":"required_actions","index$":12},"scheduled_payment_details":{"a":true,"h":"Scheduled Payment Details","n":"scheduled_payment_details","r":true,"sh":"Details of the scheduled payment","t":"`$OBJECT`","key$":"scheduled_payment_details","index$":13},"services":{"a":true,"h":"Services","n":"services","r":true,"sh":"A payment service supported by a payment provider","t":"`$ARRAY`","key$":"services","index$":14},"standing_order_details":{"a":true,"h":"Standing Order Details","n":"standing_order_details","r":true,"sh":"Details of the standing order","t":"`$OBJECT`","key$":"standing_order_details","index$":15},"supported_currencies":{"a":true,"h":"Supported Currencies","n":"supported_currencies","r":true,"sh":"The currencies supported by the payment provider","t":"`$ARRAY`","key$":"supported_currencies","index$":16}},"name":"initiate_payment_bud_license","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /v1/payments/scheduled/bud-pay-url","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/v1/payments/scheduled/bud-pay-url","q":{"exist":["x_client_id"]},"r":{},"s":[{"lit":"v1"},{"lit":"payments"},{"lit":"scheduled"},{"lit":"bud-pay-url"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":0},{"a":true,"co":{"id":"POST /v1/payments/single/bud-pay-url","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/v1/payments/single/bud-pay-url","q":{"exist":["x_client_id"]},"r":{},"s":[{"lit":"v1"},{"lit":"payments"},{"lit":"single"},{"lit":"bud-pay-url"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":1},{"a":true,"co":{"id":"POST /v1/payments/standing-order/bud-pay-url","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/v1/payments/standing-order/bud-pay-url","q":{"exist":["x_client_id"]},"r":{},"s":[{"lit":"v1"},{"lit":"payments"},{"lit":"standing-order"},{"lit":"bud-pay-url"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":2}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /v1/payments/providers","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0}],"query":[{"a":true,"ex":"GBR,DEU","k":"query","n":"country","or":"country","r":false,"t":"`$STRING`","index$":0},{"a":true,"k":"query","n":"product","or":"product","r":false,"t":"`$ARRAY`","index$":1},{"a":true,"k":"query","n":"type","or":"type","r":false,"t":"`$STRING`","index$":2}]},"k":"http","m":"GET","o":"/v1/payments/providers","q":{"exist":["country","product","type","x_client_id"]},"r":{},"s":[{"lit":"v1"},{"lit":"payments"},{"lit":"providers"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"initiate_payment_bud_license","name__orig":"initiate_payment_bud_license","Name":"InitiatePaymentBudLicense","name_":"initiate_payment_bud_license","name-":"initiate-payment-bud-license","NAME":"INITIATE_PAYMENT_BUD_LICENSE","index$":0}, {"active":true,"entity":"initiate_payment_bud_license","key$":"BasicInitiatePaymentBudLicenseFlow","kind":"basic","name":"BasicInitiatePaymentBudLicenseFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"initiate_payment_bud_license_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"initiate_payment_bud_license_ref01"}}],"index$":1}]}, 'InitiatePaymentBudLicense', {"POST /v1/payments/scheduled/bud-pay-url":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"title":"Scheduled Payment Bud Pay URL Payload","type":"object","required":["redirect_url","provider","scheduled_payment_details"],"properties":{"redirect_url":{"type":"string","description":"URL where the user will be redirected to once they have completed the TPP Connection flow. This can be a web based URL or a mobile application internal URL.","key$":"redirect_url"},"provider":{"type":"string","description":"The name (identifier) of the payment provider","key$":"provider"},"provider_types":{"title":"Provider Types","type":"array","description":"One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. If no filter is applied, all types of provider are returned. Note that on Bud's sandbox environment, only sandbox providers will be returned.","items":{"type":"string","enum":["business","retail","sandbox"]},"x-ref":"#/paths/~1v1~1payments~1scheduled/post/requestBody/content/application~1json/schema/properties/provider_types","key$":"provider_types"},"scheduled_payment_details":{"title":"Scheduled Payment Details","type":"object","additionalProperties":false,"description":"Details of the scheduled payment","required":["reference","recipient","sender","amount","requested_execution_date"],"properties":{"reference":{"type":"string","description":"Reference to use for the payment made through this scheduled instruction","maxLength":18},"recipient":{"title":"Payment Recipient","type":"object","description":"Details of the recipient of the payment","required":["type","name","account_number"],"properties":{"type":{},"name":{},"account_number":{}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details/properties/recipient"},"sender":{"title":"Payment Sender","type":"object","description":"Details of the payment sender","required":["user_id","name"],"properties":{"user_id":{},"type":{},"name":{},"account_number":{},"provider":{}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details/properties/sender"},"amount":{"title":"Standard Bud Amount","type":"object","description":"The monetary amount.","required":["value","currency"],"properties":{"value":{},"currency":{}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details/properties/amount"},"requested_execution_date":{"type":"string","format":"date-time","description":"Date-time (ISO 8601) the payment should be made. Must be at least one day in the future.","example":"2020-06-01T15:00:00Z"}},"x-ref":"#/paths/~1v1~1payments~1scheduled/post/requestBody/content/application~1json/schema/properties/scheduled_payment_details","key$":"scheduled_payment_details"}},"index$":1},"example":{"redirect_url":"https://client-redirect-url/example","provider":"natwest","provider_types":["retail"],"scheduled_payment_details":{"reference":"PAYMENT REFERENCE","recipient":{"type":"SortCodeAccountNumber","name":"MR TEST RECIPIENT","account_number":"01234501234567"},"sender":{"user_id":"client-defined-user-id","name":"MR TEST SENDER"},"amount":{"value":"2.50","currency":"GBP"},"requested_execution_date":"2023-01-20T15:00:00Z"}}}}},"parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1v1~1payments~1single/get/parameters/0","index$":0}]},"POST /v1/payments/single/bud-pay-url":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"title":"Payments Single Bud Pay URL Payload","type":"object","required":["redirect_url","provider","payment_details"],"properties":{"redirect_url":{"type":"string","description":"URL where the user will be redirected to once they have completed the TPP Connection flow. This can be a web based URL or a mobile application internal URL.","nullable":false,"maxLength":2083,"key$":"redirect_url"},"provider":{"type":"string","description":"The name (identifier) of the payment provider","key$":"provider"},"provider_types":{"title":"Provider Types","type":"array","description":"One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. If no filter is applied, all types of provider are returned. Note that on Bud's sandbox environment, only sandbox providers will be returned.","items":{"type":"string","enum":["business","retail","sandbox"]},"x-ref":"#/paths/~1v1~1payments~1scheduled/post/requestBody/content/application~1json/schema/properties/provider_types","key$":"provider_types"},"payment_details":{"title":"Single Payment Details","type":"object","description":"Payment details required to make a Single Payment","required":["reference","recipient","sender","amount"],"properties":{"reference":{"type":"string","description":"The payment reference"},"recipient":{"title":"Payment Recipient","type":"object","description":"Details of the recipient of the payment","required":["type","name","account_number"],"properties":{"type":{},"name":{},"account_number":{}}},"sender":{"title":"Payment Sender","type":"object","description":"Details of the payment sender","required":["user_id","name"],"properties":{"user_id":{},"type":{},"name":{},"account_number":{},"provider":{}}},"amount":{"title":"Standard Bud Amount","type":"object","description":"The monetary amount.","required":["value","currency"],"properties":{"value":{},"currency":{}}}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details","key$":"payment_details"}},"index$":1},"examples":{"Bud Pay URL UK":{"value":{"redirect_url":"https://client-redirect-url/example","provider":"natwest","provider_types":["retail"],"payment_details":{"reference":"PAYMENT REFERENCE","recipient":{"type":"SortCodeAccountNumber","name":"MR TEST RECIPIENT","account_number":"01234501234567"},"sender":{"user_id":"client-defined-user-id","name":"Mr TEST SENDER"},"amount":{"value":"2.50","currency":"GBP"}}}},"Bud Pay URL Europe":{"value":{"redirect_url":"https://client-redirect-url/foo","provider":"commerzbank","payment_details":{"reference":"REFERENCE FOR PAYMENT","recipient":{"type":"IBAN","name":"MR TEST RECIPIENT","account_number":"DE40100100103307118603"},"sender":{"user_id":"client-defined-user-id","name":"Mr TEST SENDER","type":"IBAN","account_number":"DE50100100103307118608"},"amount":{"value":"2.50","currency":"EUR"}}}}}}}},"parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1v1~1payments~1single/get/parameters/0","index$":0}]},"POST /v1/payments/standing-order/bud-pay-url":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"title":"Payments Standing Order Bud Pay URL Payload","type":"object","required":["redirect_url","provider","standing_order_details"],"properties":{"redirect_url":{"type":"string","description":"URL where the user will be redirected to once they have completed the TPP Connection flow. This can be a web based URL or a mobile application internal URL.","key$":"redirect_url"},"provider":{"type":"string","description":"The name (identifier) of the payment provider","key$":"provider"},"provider_types":{"title":"Provider Types","type":"array","description":"One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. If no filter is applied, all types of provider are returned. Note that on Bud's sandbox environment, only sandbox providers will be returned.","items":{"type":"string","enum":["business","retail","sandbox"]},"x-ref":"#/paths/~1v1~1payments~1scheduled/post/requestBody/content/application~1json/schema/properties/provider_types","key$":"provider_types"},"standing_order_details":{"title":"Standing Order Details","type":"object","additionalProperties":false,"description":"Details of the standing order","required":["reference","recipient","sender","recurring_amount","first_payment_date","frequency"],"properties":{"reference":{"type":"string","description":"Reference to use for payments made through this Standing Order","maxLength":18},"recipient":{"title":"Payment Recipient","type":"object","description":"Details of the recipient of the payment","required":["type","name","account_number"],"properties":{"type":{},"name":{},"account_number":{}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details/properties/recipient"},"sender":{"title":"Payment Sender","type":"object","description":"Details of the payment sender","required":["user_id","name"],"properties":{"user_id":{},"type":{},"name":{},"account_number":{},"provider":{}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details/properties/sender"},"recurring_amount":{"title":"Standard Bud Amount","type":"object","description":"The monetary amount.","required":["value","currency"],"properties":{"value":{},"currency":{}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details/properties/amount"},"first_payment_date":{"type":"string","format":"date-time","description":"Date-time (ISO 8601) the first payment should be made. Must be a minimum of 3 business days in the future and must be on a business day.","example":"2020-06-01T15:00:00Z"},"last_payment_date":{"type":"string","format":"date-time","description":"Date-time (ISO 8601) after which payments should stop. Must be after `first_payment_date`.","example":"2020-06-01T15:00:00Z"},"frequency":{"title":"Standing Order Frequency","type":"string","description":"The payment frequency for a Standing Order","enum":["daily","weekly","fortnightly","monthly","quarterly","semiannually","yearly"]}},"x-ref":"#/paths/~1v1~1payments~1standing-order/post/requestBody/content/application~1json/schema/properties/standing_order_details","key$":"standing_order_details"}},"index$":1},"example":{"redirect_url":"https://client-redirect-url/example","provider_types":["retail"],"provider":"BankOfBudOBUK","standing_order_details":{"reference":"PAYMENT REFERENCE","recipient":{"type":"SortCodeAccountNumber","name":"MR TEST RECIPIENT","account_number":"01234501234567"},"sender":{"user_id":"client-defined-user-id","name":"MR TEST SENDER"},"recurring_amount":{"value":"2.50","currency":"GBP"},"first_payment_date":"2023-02-10T15:00:00Z","last_payment_date":"2023-02-10T15:00:00Z","frequency":"monthly"}}}}},"parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1v1~1payments~1single/get/parameters/0","index$":0}]},"GET /v1/payments/providers":{"protocol":"http","parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1v1~1payments~1single/get/parameters/0","index$":0},{"in":"query","name":"country","description":"Country filter as a comma separated list of country codes following the ISO 3166 alpha-3 code format.","example":"GBR,DEU","schema":{"type":"string"},"index$":1},{"in":"query","name":"product","description":"Product filter as a comma separated list of bud supported products","schema":{"title":"Payment Services","description":"A payment service supported by a payment provider","type":"array","minItems":1,"items":{"type":"string","enum":["domestic-single-payment","domestic-standing-order","domestic-scheduled-payment"]}},"index$":2},{"in":"query","required":false,"name":"type","description":"A comma separated list of Bud support provider types to be shown. If no filter is applied, all the providers are returned.","schema":{"type":"string","enum":["sandbox","business","retail"]},"index$":3}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const initiate_payment_bud_license_ref01_ent = client.InitiatePaymentBudLicense()
    let initiate_payment_bud_license_ref01_data = setup.data.new.initiate_payment_bud_license['initiate_payment_bud_license_ref01']

    initiate_payment_bud_license_ref01_data = (await initiate_payment_bud_license_ref01_ent.create(initiate_payment_bud_license_ref01_data)).data()
    assert(null != initiate_payment_bud_license_ref01_data)


    // LIST
    const initiate_payment_bud_license_ref01_match = {}

    const initiate_payment_bud_license_ref01_list = (await initiate_payment_bud_license_ref01_ent.list(initiate_payment_bud_license_ref01_match)).map((e) => e.data())


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/initiate_payment_bud_license/InitiatePaymentBudLicenseTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = BudPaymentsSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['initiate_payment_bud_license01','initiate_payment_bud_license02','initiate_payment_bud_license03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'BUD_PAYMENTS_TEST_INITIATE_PAYMENT_BUD_LICENSE_ENTID': idmap,
    'BUD_PAYMENTS_TEST_LIVE': 'FALSE',
    'BUD_PAYMENTS_TEST_EXPLAIN': 'FALSE',
    'BUD_PAYMENTS_APIKEY': '',
  })

  idmap = env['BUD_PAYMENTS_TEST_INITIATE_PAYMENT_BUD_LICENSE_ENTID']

  const live = 'TRUE' === env.BUD_PAYMENTS_TEST_LIVE
  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['BUD_PAYMENTS_TEST_INITIATE_PAYMENT_BUD_LICENSE_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new BudPaymentsSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
        apikey: env.BUD_PAYMENTS_APIKEY,
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when
      // the last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey and
      // server values above and handed the SDK undefined.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.BUD_PAYMENTS_TEST_EXPLAIN,
    live,
    transport,
    now: Date.now(),
  }

  return setup
}
  
