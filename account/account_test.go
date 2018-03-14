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

package account

import (
	"fmt"
	"github.com/Ontology/crypto"
	"os"
	"path"
	"testing"
)

func TestClient(t *testing.T) {
	t.Log("created client start!")
	crypto.SetAlg(crypto.P256R1)
	dir := "./data/"
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		t.Log("create dir ", dir, " error: ", err)
	} else {
		t.Log("create dir ", dir, " success!")
	}
	for i := 0; i < 10000; i++ {
		p := path.Join(dir, fmt.Sprintf("wallet%d.txt", i))
		fmt.Println("client path", p)
		CreateClient(p, []byte(DefaultPin))
	}
}
