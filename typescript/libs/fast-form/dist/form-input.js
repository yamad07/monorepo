import { FormAbstructModule } from './form-abstruct-module';
export default class FormInputModule extends FormAbstructModule {
    constructor(label, value, type, key, rules) {
        super();
        this.label = label;
        this._value = value;
        this.type = type;
        this.key = key;
        this.rules = rules || '';
    }
    set value(value) {
        this._value = value;
    }
    get value() {
        return this._value;
    }
}
