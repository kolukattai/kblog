import { FC, useEffect, useMemo, useRef, useState } from "react"

import style from './editor.module.scss'
import { DropDownMenu } from "../dropdown"

type WysiwygEditorProps = {
    data: string
    onChange: (val: string) => void
    disabled?: boolean
}

export const WysiwygEditor: FC<WysiwygEditorProps> = (_props) => {
    const divRef = useRef<HTMLIFrameElement>(null);

    useEffect(() => {
        // Log the current innerHTML of the div on initial render
        console.log(divRef.current?.innerHTML);

        if (!!divRef.current && divRef.current.contentDocument && divRef.current.contentDocument.body) {
            // divRef.current.contentDocument.designMode = "on"

            const observer = new MutationObserver(mutationsList => {
                mutationsList.forEach(mutation => {
                    _props.onChange(`${divRef.current?.contentDocument?.documentElement.innerHTML}`)
                });
            });

            observer.observe(divRef.current.contentDocument, { subtree: true, childList: true, characterData: true });
    
            return () => {
                observer.disconnect();
            };
        }

        // Clean up the observer on component unmount
    }, [divRef.current]);

    const [titleList, setTitleList] = useState<{label: string, value: string}[]>([])
    const [formateList, setFormateList] = useState<{label: string, value: string}[]>([])
    
    useEffect(() => {
        let items = []
        for (let i = 1; i < 7; i++) {
            items.push({label: `Title ${i}`, value: `h${i}`})
        }
        setFormateList([
            {label: "Bold", value: "b"},
            {label: "Italic", value: "i"},
            {label: "StrikeThrough", value: "s"},
        ])
        setTitleList(items)
    }, [])

    useEffect(() => {
        let ele = divRef.current?.contentDocument
        if (!!ele) {
            ele.designMode = "on"
            ele.body.innerHTML = _props.data
            addCss("http://localhost:8080/static/css/main.css")
            console.log("stuff", divRef.current?.contentDocument);
        }
    }, [divRef.current, divRef.current?.contentDocument])

    const frameData = useMemo(() => {
        return "data:text/html,"+encodeURIComponent(_props.data)
    }, [_props.data])


    const addCss = (link: string) => {
        var cssLink = document.createElement("link");
        cssLink.href = link; 
        cssLink.rel = "stylesheet"; 
        cssLink.type = "text/css";
        let ele = divRef.current?.contentDocument
        ele?.head.appendChild(cssLink)
    }

    const execElementCmd = (element: string) => {
        let ifWindow = divRef.current?.contentDocument

        if (!!!ifWindow) return

        // Get the selected text
        let selection = ifWindow.getSelection();
        if (!!!selection) return;
        if (selection.rangeCount === 0) return;

        let range = selection.getRangeAt(0);
        let selectedText = range.toString();

        console.log(range, selectedText, selection.rangeCount)
        if (!!!selectedText) {
            selectedText = `${range.startContainer.textContent}`
            let newRange = range.cloneRange()
        }
        

        // Check if the selected text is already wrapped with h1
        let parentElement = range.commonAncestorContainer.parentElement;
        if (!!!parentElement) return
        let isWrappedWithElement = parentElement.tagName.toLowerCase() === element;

        // Toggle wrapping
        if (isWrappedWithElement) {
            // Unwrap the selected text (remove h1 tags)
            let textNode = document.createTextNode(selectedText);
            parentElement.parentNode?.replaceChild(textNode, parentElement);
        } else {
            // Wrap the selected text with h1
            let selectedElement = document.createElement(element);
            selectedElement.textContent = selectedText;
            range.deleteContents();
            range.insertNode(selectedElement);
        }

        // Clear the selection (optional)
        selection.removeAllRanges();
    }

    return (
        <div className={style.editor}>
            <div>
                {/* <DropDownMenu
                items={list}
                label="Title"
                onChange={(e) => {
                    document.execCommand(e.value);
                }}
                ></DropDownMenu> */}
                <div>
                {titleList.map((e, i) => (
                    <button onClick={() => {
                        execElementCmd(e.value)
                    }} key={i}>
                        {e.label}
                    </button>
                ))}
                </div>
                <div>
                {formateList.map((e, i) => (
                    <button className="p-2 mr-2" onClick={() => {
                        execElementCmd(e.value)
                    }} key={i}>
                        {e.label}
                    </button>
                ))}
                </div>
            </div>
            <iframe
                className={style.editorContent}
                id="editor-content" 
                style={{minHeight: "90vh"}}
                ref={divRef} 
                dangerouslySetInnerHTML={{__html: _props.data}}
            ></iframe>
        </div>
    )
}