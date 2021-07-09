import { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { Accordion, AccordionItem } from 'react-sanfona';
import axios from 'axios';

import { Checkbox } from '../../components';
import { useHistory } from 'react-router';

const useQuery = () => {
  return new URLSearchParams(useLocation().search);
};

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
  const query = useQuery().get('q');
  const sort = useQuery().get('sort');
  const from = useQuery().get('from');
  const history = useHistory();

  const [products, setProducts] = useState([]);

  useEffect(() => {
    axios
      .get('https://run.mocky.io/v3/ad458ff0-224b-4c4b-b94e-7a73ec184f06')
      .then((reply) => {
        if (reply.data.success) {
          setProducts(reply.data.data);
        }
      });
  }, []);

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

      {/*Search result*/}
      <div className="flex flex-col w-full space-y-6">
        {/*Produk & Toko Tab*/}
        <div className="flex flex-row w-full border-b border-gray-200">
          <button className="flex flex-row items-center px-5 py-3 border-b-4 border-xgreen space-x-2">
            <i className="icon-product" />
            <span className="text-sm font-bold text-xgreen">Produk</span>
          </button>
          <button className="flex flex-row items-center px-5 py-3 border-b-4 border-transparent space-x-2">
            <i className="icon-shop" />
            <span className="text-sm font-bold text-gray-400">Toko</span>
          </button>
        </div>

        {/*Search info and sorting dropdown*/}
        <div className="flex flex-row w-full">
          <div className="flex flex-row w-full items-center justify-between">
            <div className="text-xs">
              Menampilkan 1 - 60 barang dari total 12.6jt+ untuk sepatu
            </div>
            <div className="flex flex-row items-center space-x-2">
              <div className="text-xs font-bold">Urutkan:</div>
              <div className="relative w-48 py-2 pl-3 pr-3 text-left bg-white rounded-lg border border-gray-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-opacity-75 focus-visible:ring-white focus-visible:ring-offset-orange-300 focus-visible:ring-offset-2 focus-visible:border-indigo-500 sm:text-sm">
                <select
                  className="w-full focus:outline-none focus:ring-0"
                  onChange={(e) => {
                    history.push({
                      search: `?from=${from}&q=${query}&sort=${e.target.value}`,
                    });
                  }}
                >
                  <option value="default">Paling Sesuai</option>
                  <option value="ratings">Ulasan</option>
                  <option value="newest">Terbaru</option>
                  <option value="highestprice">Harga Tertinggi</option>
                  <option value="lowestprice">Harga Terendah</option>
                  <option value="lowesttco">Harga + Ongkir Terendah</option>
                </select>
              </div>
            </div>
          </div>
        </div>

        {/*Product card grid*/}
        <div className="flex flex-row w-full">
          <div className="flex flex-grow w-full flex-wrap">
            {products &&
              products.map((item, idx) => (
                <div
                  key={idx}
                  className="group flex w-1/5 px-2 py-3 justify-center"
                >
                  <div className="bg-white shadow-md cursor-pointer rounded">
                    <img
                      src={item.ProductImageURL}
                      alt="product-image"
                      className="rounded-t w-full"
                    />
                    <div className="p-3 space-y-2 rounded-b text-gray-700">
                      <div
                        className="text-xs truncate whitespace-pre-wrap"
                        style={{
                          display: '-webkit-box',
                          WebkitLineClamp: 2,
                          WebkitBoxOrient: 'vertical',
                        }}
                      >
                        {item.ProductName}
                        GREEN NAVY
                      </div>
                      <div className="text-sm font-bold">
                        Rp{item.Price.toLocaleString('id-ID')}
                      </div>
                      <div className="inline-flex px-1.5 py-1 space-x-1 bg-purple-100 rounded-md text-xs text-purple-400 transform -translate-x-2 scale-90">
                        <i className="icon-kurir" />
                        <span>
                          dari
                          <b> Rp{item.ShippingCost.toLocaleString('id-ID')}</b>
                        </span>
                      </div>

                      <div className="flex flex-row text-xs text-gray-500">
                        <div className="flex flex-row group-hover:hidden">
                          <i className="icon-power-merchant" />
                          <div>{item.SellerCity}</div>
                        </div>
                        <div className="hidden flex-row group-hover:flex">
                          <i className="icon-power-merchant" />
                          <div>{item.SellerName}</div>
                        </div>
                      </div>
                      <div className="flex flex-row text-gray-500 divide-x divide-gray-300">
                        <div className="flex flex-row text-xs space-x-1 pr-2">
                          <i className="icon-star" />
                          <span>{item.Rating}</span>
                        </div>
                        <div className="flex flex-row text-xs space-x-1 pl-2">
                          Terjual {item.TotalSales}
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              ))}
          </div>
        </div>
      </div>
    </div>
  );
};
