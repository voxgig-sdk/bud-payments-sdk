import { BudPaymentsEntityBase } from '../BudPaymentsEntityBase';
import type { BudPaymentsSDK } from '../BudPaymentsSDK';
import type { Control } from '../types';
import type { ManagePayment, ManagePaymentLoadMatch, ManagePaymentListMatch, ManagePaymentCreateData } from '../BudPaymentsTypes';
declare class ManagePaymentEntity extends BudPaymentsEntityBase<ManagePayment> {
    constructor(client: BudPaymentsSDK, entopts: any);
    make(this: ManagePaymentEntity): ManagePaymentEntity;
    load(this: any, reqmatch?: ManagePaymentLoadMatch, ctrl?: Control): Promise<ManagePaymentEntity>;
    list(this: any, reqmatch?: ManagePaymentListMatch, ctrl?: Control): Promise<ManagePaymentEntity[]>;
    create(this: any, reqdata?: ManagePaymentCreateData, ctrl?: Control): Promise<ManagePaymentEntity>;
}
export { ManagePaymentEntity };
