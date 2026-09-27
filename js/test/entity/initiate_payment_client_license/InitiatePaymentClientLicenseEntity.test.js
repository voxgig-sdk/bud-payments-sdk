
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


describe('InitiatePaymentClientLicenseEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when BUD_PAYMENTS_TEST_LIVE=TRUE.
  afterEach(liveDelay('BUD_PAYMENTS_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = BudPaymentsSDK.test()
    const ent = testsdk.InitiatePaymentClientLicense()
    assert(null != ent)
  })


  test('basic', async (t) => {

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"authorisation_url":{"a":true,"h":"Authorisation Url","n":"authorisation_url","r":true,"sh":"The authorisation URL for a relevant payment provider.","t":"`$STRING`","key$":"authorisation_url","index$":0},"code":{"a":true,"h":"Code","n":"code","r":false,"sh":"Code parameter returned in redirect parameters","t":"`$STRING`","key$":"code","index$":1},"payment_details":{"a":true,"h":"Payment Details","n":"payment_details","r":true,"sh":"Payment details required to make a Single Payment","t":"`$OBJECT`","key$":"payment_details","index$":2},"payment_id":{"a":true,"h":"Payment Id","n":"payment_id","op":{"create":{"req":true,"type":"`$STRING`"}},"r":false,"sh":"Payment Identifier","t":"`$STRING`","key$":"payment_id","index$":3},"payment_type":{"a":true,"h":"Payment Type","n":"payment_type","r":false,"sh":"The type of payment iniated through Bud's Payments Service","t":"`$STRING`","key$":"payment_type","index$":4},"provider":{"a":true,"h":"Provider","n":"provider","r":true,"sh":"The name (identifier) of the payment provider","t":"`$STRING`","key$":"provider","index$":5},"provider_redirect_url":{"a":true,"h":"Provider Redirect Url","n":"provider_redirect_url","r":false,"sh":"This is the url that your Customer will be redirected to once they have authorised with the relevant provider.","t":"`$STRING`","key$":"provider_redirect_url","index$":6},"provider_types":{"a":true,"h":"Provider Types","n":"provider_types","r":false,"sh":"One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow.","t":"`$ARRAY`","key$":"provider_types","index$":7},"redirect_url":{"a":true,"h":"Redirect Url","n":"redirect_url","r":false,"sh":"This is the url that your Customer will be redirected to when using the Bud user interface.","t":"`$STRING`","key$":"redirect_url","index$":8},"required_action":{"a":true,"h":"Required Action","n":"required_action","r":false,"sh":"If present, this fild will indicate the next required action (if any) to be taken by the client in order to complete the payment flow | Value | Description | |:-------|:-------------| | confirm_scheduled_payment | Indicates that the client…","t":"`$STRING`","key$":"required_action","index$":9},"result":{"a":true,"h":"Result","n":"result","r":false,"t":"`$STRING`","key$":"result","index$":10},"scheduled_payment_details":{"a":true,"h":"Scheduled Payment Details","n":"scheduled_payment_details","r":true,"sh":"Details of the scheduled payment","t":"`$OBJECT`","key$":"scheduled_payment_details","index$":11},"standing_order_details":{"a":true,"h":"Standing Order Details","n":"standing_order_details","r":true,"sh":"Details of the standing order","t":"`$OBJECT`","key$":"standing_order_details","index$":12},"state":{"a":true,"h":"State","n":"state","r":true,"sh":"The task ID","t":"`$STRING`","key$":"state","index$":13}},"name":"initiate_payment_client_license","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /v1/payments/authorisation-codes","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/v1/payments/authorisation-codes","q":{"exist":["x_client_id"]},"r":{},"s":[{"lit":"v1"},{"lit":"payments"},{"lit":"authorisation-codes"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":0},{"a":true,"co":{"id":"POST /v1/payments/scheduled","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/v1/payments/scheduled","q":{"exist":["x_client_id"]},"r":{},"s":[{"lit":"v1"},{"lit":"payments"},{"lit":"scheduled"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":1},{"a":true,"co":{"id":"POST /v1/payments/single","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/v1/payments/single","q":{"exist":["x_client_id"]},"r":{},"s":[{"lit":"v1"},{"lit":"payments"},{"lit":"single"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":2},{"a":true,"co":{"id":"POST /v1/payments/standing-order","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/v1/payments/standing-order","q":{"exist":["x_client_id"]},"r":{},"s":[{"lit":"v1"},{"lit":"payments"},{"lit":"standing-order"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":3}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"initiate_payment_client_license","name__orig":"initiate_payment_client_license","Name":"InitiatePaymentClientLicense","name_":"initiate_payment_client_license","name-":"initiate-payment-client-license","NAME":"INITIATE_PAYMENT_CLIENT_LICENSE","index$":1}, {"active":true,"entity":"initiate_payment_client_license","key$":"BasicInitiatePaymentClientLicenseFlow","kind":"basic","name":"BasicInitiatePaymentClientLicenseFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"initiate_payment_client_license_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0}]}, 'InitiatePaymentClientLicense', {"POST /v1/payments/authorisation-codes":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"title":"Authorisation Codes Request Payload","type":"object","required":["state"],"additionalProperties":true,"properties":{"state":{"type":"string","description":"The task ID","key$":"state"},"code":{"type":"string","description":"Code parameter returned in redirect parameters","key$":"code"}},"index$":1},"examples":{"Success":{"value":{"state":"e017bd1c-9b39-40cc-b9ec-e21fce26719e","code":"803A940E-2416-4A34-BF5F-4AC8C9977AC4"}},"Error":{"value":{"state":"e017bd1c-9b39-40cc-b9ec-e21fce26719e","error":"access_denied"}}}}}},"parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1v1~1payments~1single/get/parameters/0","index$":0}]},"POST /v1/payments/scheduled":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"title":"Create Scheduled Payment Request","type":"object","required":["provider","scheduled_payment_details"],"properties":{"redirect_url":{"type":"string","description":"This is the url that your Customer will be redirected to when using the Bud user interface. If you wish to use your own license and callback infrastructure this parameter will have no effect and the provider_redirect_url parameter should be used instead.","key$":"redirect_url"},"provider_redirect_url":{"type":"string","description":"This is the url that your Customer will be redirected to once they have authorised with the relevant provider. This should only be used when using your own license and wish to use the Submit Authorisation Codes endpoint. If you wish to use Bud's call back then the redirect_url must be provided.","key$":"provider_redirect_url"},"provider":{"type":"string","description":"The name (identifier) of the payment provider","key$":"provider"},"provider_types":{"title":"Provider Types","type":"array","description":"One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. If no filter is applied, all types of provider are returned. Note that on Bud's sandbox environment, only sandbox providers will be returned.","items":{"type":"string","enum":["business","retail","sandbox"]},"key$":"provider_types"},"scheduled_payment_details":{"title":"Scheduled Payment Details","type":"object","additionalProperties":false,"description":"Details of the scheduled payment","required":["reference","recipient","sender","amount","requested_execution_date"],"properties":{"reference":{"type":"string","description":"Reference to use for the payment made through this scheduled instruction","maxLength":18},"recipient":{"title":"Payment Recipient","type":"object","description":"Details of the recipient of the payment","required":["type","name","account_number"],"properties":{"type":{},"name":{},"account_number":{}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details/properties/recipient"},"sender":{"title":"Payment Sender","type":"object","description":"Details of the payment sender","required":["user_id","name"],"properties":{"user_id":{},"type":{},"name":{},"account_number":{},"provider":{}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details/properties/sender"},"amount":{"title":"Standard Bud Amount","type":"object","description":"The monetary amount.","required":["value","currency"],"properties":{"value":{},"currency":{}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details/properties/amount"},"requested_execution_date":{"type":"string","format":"date-time","description":"Date-time (ISO 8601) the payment should be made. Must be at least one day in the future.","example":"2020-06-01T15:00:00Z"}},"key$":"scheduled_payment_details"}},"index$":1},"example":{"redirect_url":"https://your-app.com/connected?my-parameter=xyz","provider":"lloyds","scheduled_payment_details":{"reference":"PAYMENT REFERENCE","recipient":{"type":"SortCodeAccountNumber","name":"MR TEST RECIPIENT","account_number":"01234501234567"},"sender":{"user_id":"client-defined-user-id","name":"MR TEST SENDER"},"amount":{"value":"2.50","currency":"GBP"},"requested_execution_date":"2020-02-02T12:29:33.428944815Z"}}}}},"parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1v1~1payments~1single/get/parameters/0","index$":0}]},"POST /v1/payments/single":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"title":"Single Payment Request Payload","type":"object","description":"Expected request payload structure","required":["provider","payment_details"],"properties":{"redirect_url":{"type":"string","description":"This is the url that your Customer will be redirected to when using the Bud user interface. If you wish to use your own license and callback infrastructure this parameter will have no effect and the provider_redirect_url parameter should be used instead.","key$":"redirect_url"},"provider_redirect_url":{"type":"string","description":"This is the url that your Customer will be redirected to once they have authorised with the relevant provider. This should only be used when using your own license and wish to use the Submit Authorisation Codes endpoint. If you wish to use Bud's call back then the redirect_url must be provided.","key$":"provider_redirect_url"},"provider":{"type":"string","description":"The name (identifier) of the payment provider","key$":"provider"},"payment_details":{"title":"Single Payment Details","type":"object","description":"Payment details required to make a Single Payment","required":["reference","recipient","sender","amount"],"properties":{"reference":{"type":"string","description":"The payment reference"},"recipient":{"title":"Payment Recipient","type":"object","description":"Details of the recipient of the payment","required":["type","name","account_number"],"properties":{"type":{},"name":{},"account_number":{}}},"sender":{"title":"Payment Sender","type":"object","description":"Details of the payment sender","required":["user_id","name"],"properties":{"user_id":{},"type":{},"name":{},"account_number":{},"provider":{}}},"amount":{"title":"Standard Bud Amount","type":"object","description":"The monetary amount.","required":["value","currency"],"properties":{"value":{},"currency":{}}}},"key$":"payment_details"}},"index$":1},"examples":{"Create Payment UK":{"value":{"redirect_url":"https://your-app.com/connected?my-parameter=xyz","provider":"lloyds","payment_details":{"reference":"PAYMENT REFERENCE","recipient":{"type":"SortCodeAccountNumber","name":"MR TEST RECIPIENT","account_number":"01234501234567"},"sender":{"user_id":"client-defined-user-id","name":"MR TEST SENDER"},"amount":{"value":"2.50","currency":"GBP"}}}},"Create Payment Europe":{"value":{"redirect_url":"https://your-app.com/connected?my-parameter=xyz","provider":"commerzbank","payment_details":{"reference":"PAYMENT REFERENCE","recipient":{"type":"IBAN","name":"MR TEST RECIPIENT","account_number":"DE40100100103307118605"},"sender":{"user_id":"client-defined-user-id","type":"IBAN","name":"MR TEST SENDER","account_number":"DE02100100109307118608"},"amount":{"value":"2.50","currency":"EUR"}}}}}}}},"parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1v1~1payments~1single/get/parameters/0","index$":0}]},"POST /v1/payments/standing-order":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"title":"Create Standing Order Request","type":"object","required":["provider","standing_order_details"],"properties":{"redirect_url":{"type":"string","description":"This is the url that your Customer will be redirected to when using the Bud user interface. If you wish to use your own license and callback infrastructure this parameter will have no effect and the provider_redirect_url parameter should be used instead.","key$":"redirect_url"},"provider_redirect_url":{"type":"string","description":"This is the url that your Customer will be redirected to once they have authorised with the relevant provider. This should only be used when using your own license and wish to use the Submit Authorisation Codes endpoint. If you wish to use Bud's call back then the redirect_url must be provided.","key$":"provider_redirect_url"},"provider":{"type":"string","description":"The name (identifier) of the payment provider","key$":"provider"},"provider_types":{"title":"Provider Types","type":"array","description":"One or more provider types that can be used to filter the list of providers that are shown to the customer during the connection flow. If no filter is applied, all types of provider are returned. Note that on Bud's sandbox environment, only sandbox providers will be returned.","items":{"type":"string","enum":["business","retail","sandbox"]},"x-ref":"#/paths/~1v1~1payments~1scheduled/post/requestBody/content/application~1json/schema/properties/provider_types","key$":"provider_types"},"standing_order_details":{"title":"Standing Order Details","type":"object","additionalProperties":false,"description":"Details of the standing order","required":["reference","recipient","sender","recurring_amount","first_payment_date","frequency"],"properties":{"reference":{"type":"string","description":"Reference to use for payments made through this Standing Order","maxLength":18},"recipient":{"title":"Payment Recipient","type":"object","description":"Details of the recipient of the payment","required":["type","name","account_number"],"properties":{"type":{},"name":{},"account_number":{}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details/properties/recipient"},"sender":{"title":"Payment Sender","type":"object","description":"Details of the payment sender","required":["user_id","name"],"properties":{"user_id":{},"type":{},"name":{},"account_number":{},"provider":{}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details/properties/sender"},"recurring_amount":{"title":"Standard Bud Amount","type":"object","description":"The monetary amount.","required":["value","currency"],"properties":{"value":{},"currency":{}},"x-ref":"#/paths/~1v1~1payments~1single/post/requestBody/content/application~1json/schema/properties/payment_details/properties/amount"},"first_payment_date":{"type":"string","format":"date-time","description":"Date-time (ISO 8601) the first payment should be made. Must be a minimum of 3 business days in the future and must be on a business day.","example":"2020-06-01T15:00:00Z"},"last_payment_date":{"type":"string","format":"date-time","description":"Date-time (ISO 8601) after which payments should stop. Must be after `first_payment_date`.","example":"2020-06-01T15:00:00Z"},"frequency":{"title":"Standing Order Frequency","type":"string","description":"The payment frequency for a Standing Order","enum":["daily","weekly","fortnightly","monthly","quarterly","semiannually","yearly"]}},"key$":"standing_order_details"}},"index$":1},"example":{"redirect_url":"https://your-app.com/connected?my-parameter=xyz","provider":"lloyds","standing_order_details":{"reference":"PAYMENT REFERENCE","recipient":{"type":"SortCodeAccountNumber","name":"MR TEST RECIPIENT","account_number":"01234501234567"},"sender":{"user_id":"client-defined-user-id","name":"MR TEST SENDER"},"recurring_amount":{"value":"2.50","currency":"GBP"},"first_payment_date":"2020-02-02T12:29:33.428944815Z","last_payment_date":"2021-02-02T12:29:33.428944815Z","frequency":"monthly"}}}}},"parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1v1~1payments~1single/get/parameters/0","index$":0}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const initiate_payment_client_license_ref01_ent = client.InitiatePaymentClientLicense()
    let initiate_payment_client_license_ref01_data = setup.data.new.initiate_payment_client_license['initiate_payment_client_license_ref01']

    initiate_payment_client_license_ref01_data = (await initiate_payment_client_license_ref01_ent.create(initiate_payment_client_license_ref01_data)).data()
    assert(null != initiate_payment_client_license_ref01_data)


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/initiate_payment_client_license/InitiatePaymentClientLicenseTestData.json')

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
    ['initiate_payment_client_license01','initiate_payment_client_license02','initiate_payment_client_license03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'BUD_PAYMENTS_TEST_INITIATE_PAYMENT_CLIENT_LICENSE_ENTID': idmap,
    'BUD_PAYMENTS_TEST_LIVE': 'FALSE',
    'BUD_PAYMENTS_TEST_EXPLAIN': 'FALSE',
    'BUD_PAYMENTS_APIKEY': '',
  })

  idmap = env['BUD_PAYMENTS_TEST_INITIATE_PAYMENT_CLIENT_LICENSE_ENTID']

  const live = 'TRUE' === env.BUD_PAYMENTS_TEST_LIVE
  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['BUD_PAYMENTS_TEST_INITIATE_PAYMENT_CLIENT_LICENSE_ENTID']
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
  
