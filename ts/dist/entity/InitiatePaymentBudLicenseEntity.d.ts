import { BudPaymentsEntityBase } from '../BudPaymentsEntityBase';
import type { BudPaymentsSDK } from '../BudPaymentsSDK';
import type { Control } from '../types';
import type { InitiatePaymentBudLicense, InitiatePaymentBudLicenseListMatch, InitiatePaymentBudLicenseCreateData } from '../BudPaymentsTypes';
declare class InitiatePaymentBudLicenseEntity extends BudPaymentsEntityBase<InitiatePaymentBudLicense> {
    constructor(client: BudPaymentsSDK, entopts: any);
    make(this: InitiatePaymentBudLicenseEntity): InitiatePaymentBudLicenseEntity;
    list(this: any, reqmatch?: InitiatePaymentBudLicenseListMatch, ctrl?: Control): Promise<InitiatePaymentBudLicenseEntity[]>;
    create(this: any, reqdata?: InitiatePaymentBudLicenseCreateData, ctrl?: Control): Promise<InitiatePaymentBudLicenseEntity>;
}
export { InitiatePaymentBudLicenseEntity };
