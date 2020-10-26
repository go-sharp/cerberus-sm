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
    };
}
