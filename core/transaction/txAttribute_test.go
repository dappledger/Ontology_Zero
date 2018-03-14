/*
 * Copyright (C) 2018 Onchain <onchain@onchain.com>
 *
 * This file is part of The ontology_Zero.
 *
 * The ontology_Zero is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology_Zero is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology_Zero.  If not, see <http://www.gnu.org/licenses/>.
 */

package transaction

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTxAttribute(t *testing.T) {
	b := new(bytes.Buffer)

	tx := NewTxAttribute(DescriptionUrl, []byte("http:\\www.onchain.com"))
	tx.Serialize(b)
	fmt.Println("Serialize complete")

	tm := TxAttribute{DescriptionUrl, nil, 0}
	tm.Deserialize(b)
	fmt.Println("Deserialize complete.")

	fmt.Printf("Print: Usage= :0x%x,Url Date: %q\n", tm.Usage, tm.Date)
}
