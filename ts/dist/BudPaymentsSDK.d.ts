import { InitiatePaymentBudLicenseEntity } from './entity/InitiatePaymentBudLicenseEntity';
import { InitiatePaymentClientLicenseEntity } from './entity/InitiatePaymentClientLicenseEntity';
import { ManagePaymentEntity } from './entity/ManagePaymentEntity';
export type * from './BudPaymentsTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { BudPaymentsEntityBase } from './BudPaymentsEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class BudPaymentsSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    InitiatePaymentBudLicense(entopts?: Record<string, any>): InitiatePaymentBudLicenseEntity;
    InitiatePaymentClientLicense(entopts?: Record<string, any>): InitiatePaymentClientLicenseEntity;
    ManagePayment(entopts?: Record<string, any>): ManagePaymentEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): BudPaymentsSDK;
    tester(testopts?: any, sdkopts?: any): BudPaymentsSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof BudPaymentsSDK;
export { stdutil, config, BaseFeature, BudPaymentsEntityBase, BudPaymentsSDK, SDK, };
