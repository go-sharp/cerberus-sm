
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
        type
    };
}

export { isLoading, createMsg };