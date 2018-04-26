/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except 
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and 
 * limitations under the License.
 */
 
package model

// check the interface
var _ Attribute = (*attribute)(nil)

// attribute the metadata structure definition of the model attribute
type attribute struct {
	OwnerID       string `json:"bk_supplier_account"`
	ObjectID      string `json:"bk_obj_id"`
	PropertyID    string `json:"bk_property_id"`
	PropertyName  string `json:"bk_property_name"`
	PropertyGroup string `json:"bk_property_group"`
	PropertyIndex int    `json:"bk_property_index"`
	Unit          string `json:"unit"`
	Placeholder   string `json:"placeholder"`
	IsEditable    bool   `json:"editable"`
	IsPre         bool   `json:"ispre"`
	IsRequired    bool   `json:"isrequired"`
	IsReadOnly    bool   `json:"isreadonly"`
	IsOnly        bool   `json:"isonly"`
	IsSystem      bool   `json:"bk_issystem"`
	IsAPI         bool   `json:"bk_isapi"`
	PropertyType  string `json:"bk_property_type"`
	Option        string `json:"option"`
	Description   string `json:"description"`
	Creator       string `json:"creator"`
}

func (cli *attribute) Save() error {
	return nil
}

func (cli *attribute) SetID(id string) {
	cli.PropertyID = id
}

func (cli *attribute) SetName(name string) {
	cli.PropertyName = name
}

func (cli *attribute) SetUnit(unit string) {
	cli.Unit = unit
}

func (cli *attribute) SetPlaceholer(placeHoler string) {
	cli.Placeholder = placeHoler
}

func (cli *attribute) SetEditable() {
	cli.IsEditable = true
}

func (cli *attribute) SetNonEditable() {
	cli.IsEditable = false
}

func (cli *attribute) Editable() bool {
	return cli.IsEditable
}

func (cli *attribute) SetRequired() {
	cli.IsRequired = true
}

func (cli *attribute) SetNonRequired() {
	cli.IsRequired = false
}

func (cli *attribute) Required() bool {
	return cli.IsRequired
}

func (cli *attribute) SetKey(isKey bool) {
	cli.IsOnly = isKey
}

func (cli *attribute) SetOption(option string) {
	cli.Option = option
}

func (cli *attribute) SetDescrition(des string) {
	cli.Description = des
}
