import React from 'react'
import { Button, Header, Image, Modal } from 'semantic-ui-react'
import { coinbaseTicker } from '../struct'
import { StoreContainer } from '../store';

interface ProductDetailsProps {
    data: coinbaseTicker
}

const ProductDetails: React.FC<ProductDetailsProps> = () => {
    const store = StoreContainer.useContainer()
    return (
        <div>
            <p>sdhkgfuwegfitg</p>>
        </div>
    //    <Modal>
    //          {/* {store.ticker?store.ticker.map((data:coinbaseTicker)=>{ */}
    //                 {/* return ( */}
            
    //          <Modal.Header>this is a header</Modal.Header>
    //         <Modal.Content>
    //         <Modal.Description>
    //         <p>
    //         We've found the following gravatar image associated with your e-mail
    //         address.
    //         </p>
    //         <p>Is it okay to use this photo?</p>
    //         </Modal.Description>
    //         </Modal.Content>
            
    //          {/* )    */}
    //     {/* }):null} */}
    //      </Modal>  
       
    )
}


    
    
export default ProductDetails
