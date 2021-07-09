import { Accordion, AccordionItem } from 'react-sanfona';
import { Checkbox, Dropdown } from '../../components';

const Item = ({ title, children, ...rest }) => (
  <AccordionItem
    bodyClassName="flex flex-row cursor-pointer content-center justify-between"
    title={
      <div className="flex flex-row cursor-pointer content-center justify-between p-2 border border-r-0 border-l-0 border-gray-200">
        <span className="text-xs font-bold font-nunito text-gray-700">
          {title}
        </span>
        <i className="icon-navbar arrowdown" />
      </div>
    }
    {...rest}
  >
    <div className="flex flex-col w-full px-3 py-3 text-xs">{children}</div>
  </AccordionItem>
);

export const Search = () => {
  return (
    <div className="flex flex-row flex-grow w-full h-full px-12 py-10 space-x-10 bg-white">
      {/*Search filters*/}
      <div className="flex flex-col w-56">
        <span className="text-sm font-nunito font-bold">Filter</span>

        <Accordion allowMultiple className="flex flex-col my-5 shadow-md">
          <Item title="Kategori">No categories yet!</Item>
          <Item title="Jenis Toko" expanded>
            <Checkbox label="Official Store" />
            <Checkbox label="Power Merchant Pro" />
            <Checkbox label="Power Merchant" />
          </Item>
          <Item title="Lokasi" expanded>
            <Checkbox label="Palembang" />
            <Checkbox label="Surabaya" />
            <Checkbox label="Semarang" />
            <Checkbox label="DKI Jakarta" />
            <Checkbox label="Denpasar" />
            <span className="text-xgreen mt-2">Lihat selengkapnya</span>
          </Item>
          <Item title="Kondisi" expanded>
            <Checkbox label="Baru" />
            <Checkbox label="Bekass" />
          </Item>
          <Item title="Rating" expanded>
            <Checkbox
              label={
                <>
                  <i className="icon-star" />
                  <span>4 Keatas</span>
                </>
              }
            />
          </Item>
          <Item title="Penawaran" expanded>
            <Checkbox label="Cashback" />
            <Checkbox label="Gratis Ongkir" />
            <Checkbox label="COD" />
            <Checkbox label="Diskon" />
            <Checkbox label="Harga Grosir" />
          </Item>
          <Item title="Lainnya" expanded>
            <Checkbox label="Pre Order" />
            <Checkbox label="Ready Stock" />
          </Item>
        </Accordion>
      </div>
    </div>
  );
};
