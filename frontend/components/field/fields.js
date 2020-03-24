import React, { useState } from "react";
import Form from "react-bootstrap/Form";

/**
 * field is a general field used in context of authentication
 * for the application.
 */
class field {
    /**
     * @param {string} label - Field label
     * @param {string} value - initial field value
     * @param {string} type - input type
     * @param {string} placeholder - input placeholder
     * @param {RegExp} pattern - validation regex
     * @param {string} msg - message shown when invalid
     */
    constructor(label, value, type, placeholder, pattern, msg) {
        // using hooks to invoke re-render
        let [val, setVal] = useState(value);
        let [err, setErr] = useState("");
        this.value = val;
        this.err = err;
        this.setVal = setVal;
        this.setErr = setErr;

        this.label = label;
        this.type = type;
        this.placeholder = placeholder;
        this.pattern = pattern;
        this.msg = msg;
    }

    /**
     * set changes the value of the field to the one provided.
     * @param {*} newVal - new value for the field
     */
    set(newVal) {
        this.value = newVal;
        this.setVal(newVal);
    }

    validate() {
        if (!this.pattern.test(this.value)) {
            this.setErr(this.msg);
            return false;
        } else {
            this.setErr("");
            return true;
        }
    }

    paint() {
        return (
            <Form.Group>
                <Form.Label>{this.label}</Form.Label>
                <Form.Control onChange={(e) => this.setVal(e.target.value)}
                    value={this.value}
                    type={this.type}
                    placeholder={this.placeholder}
                />
                <div className="field-error">
                    {this.err}
                </div>
            </Form.Group>
        )
    }
}

export default field;
