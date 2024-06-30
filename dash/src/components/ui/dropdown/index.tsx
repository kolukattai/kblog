import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/react'
import { ChevronDownIcon } from '@heroicons/react/20/solid'
import { FC } from 'react'

function classNames(...classes: any) {
  return classes.filter(Boolean).join(' ')
}

type DropDownProps = {
    items: {
        label: string,
        value: string,
        icons: any,
    }[]
    label: string
    onChange: (val: any) => void
}

export const DropDownMenu: FC<DropDownProps> = (_props) => {
  return (
    <div>
    <Menu as="div" className="relative inline-block text-left">
      <div>
        <MenuButton className="inline-flex w-full justify-center gap-x-1.5 rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50">
          {_props.label}
          <ChevronDownIcon className="-mr-1 h-5 w-5 text-gray-400" aria-hidden="true" />
        </MenuButton>
      </div>

      <MenuItems
        transition
        className="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 transition focus:outline-none data-[closed]:scale-95 data-[closed]:transform data-[closed]:opacity-0 data-[enter]:duration-100 data-[leave]:duration-75 data-[enter]:ease-out data-[leave]:ease-in"
      >
        <div className="py-1">
            {_props.items.map((e, i) => (
                <MenuItem key={i}>
                    {({ focus }) => (
                    <button
                        onClick={() => _props.onChange(e)}
                        className={classNames(focus ? 'bg-gray-100 text-gray-900' : 'text-gray-700', 'block px-4 py-2 text-sm w-full')}
                    >
                        {e.label}
                    </button>
                    )}
                </MenuItem>
            ))}
        </div>
      </MenuItems>
    </Menu>
    </div>
  )
}
