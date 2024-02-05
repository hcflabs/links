import {Formik, Field, Form, FormikHelpers} from "formik";
import LinkForm from "@/components/LinkForm";

export default function LinkEdit({params}: { params: { slug: string } }) {
    // const {slug} = params;
    return < LinkForm owner={params.slug} short_url={params.slug} target_url={""}/>

}
