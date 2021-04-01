/**
 * Creates a new EditBaseModel.
 */
export function createEditBaseModel() {
    return {
        name: '',
        display_name: '',
        description: '',
        exe_path: '',
        work_dir: '',
        args: [],
        env: [],

        recovery_actions: [],
        stop_signal: 0,

        dependencies: [],
        service_user: '',
        start_type: 0
    };
}
