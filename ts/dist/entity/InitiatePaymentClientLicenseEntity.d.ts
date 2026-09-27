import { BudPaymentsEntityBase } from '../BudPaymentsEntityBase';
import type { BudPaymentsSDK } from '../BudPaymentsSDK';
import type { Control } from '../types';
import type { InitiatePaymentClientLicense, InitiatePaymentClientLicenseCreateData } from '../BudPaymentsTypes';
declare class InitiatePaymentClientLicenseEntity extends BudPaymentsEntityBase<InitiatePaymentClientLicense> {
    constructor(client: BudPaymentsSDK, entopts: any);
    make(this: InitiatePaymentClientLicenseEntity): InitiatePaymentClientLicenseEntity;
    create(this: any, reqdata?: InitiatePaymentClientLicenseCreateData, ctrl?: Control): Promise<InitiatePaymentClientLicenseEntity>;
}
export { InitiatePaymentClientLicenseEntity };
