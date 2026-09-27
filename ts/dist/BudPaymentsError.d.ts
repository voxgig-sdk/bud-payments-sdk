import { Context } from './Context';
declare class BudPaymentsError extends Error {
    isBudPaymentsError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { BudPaymentsError };
