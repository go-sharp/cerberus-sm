export const LOADING_EVENT = 'loading';

/**
 * Raises the loading event, to display a loading animation.
 * @param {Boolan} loading - True if loading, otherwise false.
 */
function isLoading(loading) {
    window.wails.Events.Emit(LOADING_EVENT, loading);
}

/**
 * Creates a message for Buefy toasts.
 * @param {string} msg - Message to display
 * @param {string} [typ=info] - Typ of message (info|error|warn|success)
 * @param {number} [duration=3000] - Duration of the message
 */
function createMsg(msg, typ, duration) {
    typ = typ || 'info';
    duration = duration || 3000;

    let title = 'Info';
    let type = 'is-info';
    switch (typ) {
        case 'warn':
            title = 'Warning';
            type = 'is-warning';
            break;
        case 'error':
            title = 'Error';
            type = 'is-danger';
            break;
        case 'success':
            title = 'Success';
            type = 'is-success';
            break;
        default:
            break;
    }

    return {
        duration,
        message: `<b>${title}</b>: ${msg}`,
        type,
    };
}

/**
 * Simple Debouncer to defer execution of the passed function.
 * The 'this' keyword is bound to the component passed to the constructor.
 */
class Debouncer {
    constructor() {
        this._keys = {};
    }

    /**
     * Defers execution of the passed function.
     * @param {string} key - Unique key for the passed function.
     * @param {function} fn - The function that will be executed after the defined delay.
     * @param {number} [delay=500] - Delay for the execution of the passed function.
     */
    delay(key, fn, delay = 500) {
        if (typeof fn !== 'function')
            throw new Error('Debouncer.delay: only functions allowed for fn.');
        const id = this._keys[key];
        if (id) clearInterval(id);

        this._keys[key] = setTimeout(fn, delay);
    }
}

/**
 * Checks if two objects are the same. 
 * @param {*} a - Instance a. 
 * @param {*} b - Instance b.
 */
function isEqualObj(a, b) {
    if (typeof a !== typeof b) return false;

    let same = true;
    if (Array.isArray(a)) {
        if (a.length !== b.length) return false;
        for (const idx in a) {
            same = same && isEqualObj(a[idx], b[idx]);
        }
        return same;
    } 

    if (typeof a === 'object') {
        for (const key in a) {
            same = same && isEqualObj(a[key], b[key]);
        }

        return same;
    }

    return a === b;
}

export { isLoading, createMsg, Debouncer, isEqualObj };
