import {Field, Form, Formik, FormikHelpers} from "formik";

interface Values {
    short_url: string
    target_url: string;
    owner: string;
    permanent?: boolean;
}

export default function LinkForm(params: Values) {
    return (
        <div>
            <h1>Signup</h1>
            <Formik
                initialValues={params}
                onSubmit={(
                    values: Values,
                    {setSubmitting}: FormikHelpers<Values>
                ) => {
                    setTimeout(() => {
                        values["short_url"] = params.short_url;
                        alert(JSON.stringify(values, null, 2));
                        setSubmitting(false);
                    }, 500);
                }}
            >
                <Form>
                    <div className="mt-10 grid justify-center">
                        <div className="sm:col-span-2">
                            <div className="sm:row-span-2">
                                <h1>/{params.short_url}</h1>
                            </div>

                            <div className="mt-2">
                                <div
                                    className="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md">
                                    <label
                                        htmlFor="target_url"
                                        className="block text-sm font-medium leading-6 text-gray-900"
                                    >
                                        Target Url
                                    </label>
                                    <Field id="target_url" name="target_url" placeholder={params.short_url}/>
                                </div>
                            </div>

                            <div className="mt-2">
                                <div
                                    className="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md">
                                    <label
                                        htmlFor="owner"
                                        className="block text-sm font-medium leading-6 text-gray-900"
                                    >
                                        Owner
                                    </label>
                                    <Field id="owner" name="owner" placeholder={params.short_url}/>
                                </div>
                            </div>
                            <div className="mt-2">
                                <div
                                    className="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md">
                                    <label
                                        htmlFor="owner"
                                        className="block text-sm font-medium leading-6 text-gray-900"
                                    >
                                        Permanent
                                    </label>
                                    <Field type="checkbox" name="permanent" placeholder={params.short_url}/>
                                </div>
                            </div>
                        </div>

                        <button type="submit">Submit</button>
                    </div>
                </Form>
            </Formik>
        </div>
    );
}
