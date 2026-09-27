-- BudPayments SDK error

local BudPaymentsError = {}
BudPaymentsError.__index = BudPaymentsError


function BudPaymentsError.new(code, msg, ctx)
  local self = setmetatable({}, BudPaymentsError)
  self.is_sdk_error = true
  self.sdk = "BudPayments"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function BudPaymentsError:error()
  return self.msg
end


function BudPaymentsError:__tostring()
  return self.msg
end


return BudPaymentsError
