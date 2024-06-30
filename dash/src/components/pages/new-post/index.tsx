import dynamic from 'next/dynamic';
import 'react-draft-wysiwyg/dist/react-draft-wysiwyg.css';
import { useState } from 'react';
import { WysiwygEditor } from '@/components/ui/editor';




export const NewPost = () => {
  const [data, setData] = useState("<h1>this is</h1> some shit <br/> this is other kind of shit");

  return (
    <div>
      <WysiwygEditor
        data={data}
        onChange={setData}
      />

      <div  dangerouslySetInnerHTML={{__html: data}}></div>
    </div>
  );
};



