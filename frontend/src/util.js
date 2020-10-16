
export const LOADING = 'loading';

/**
 * Raises the loading event, to display a loading animation.
 * @param {Boolan} loading - True if loading, otherwise false.
 */
function isLoading(loading) {
    window.wails.Events.Emit(LOADING, loading);
}

export  { isLoading };