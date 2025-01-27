/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mproto

import (
	"magma/orc8r/cloud/go/mproto/protos"

	"github.com/golang/protobuf/proto"
)

// MarshalManyDeterministic deterministically encodes a slice of protobuf messages,
// indexed by a unique identifier string, utilizing MarshalDeterministic.
func MarshalManyDeterministic(protosByID map[string]proto.Message) ([]byte, error) {
	// The function individually serializes each proto in order to standardize the data
	// into a uniform format to fit into a generic map<string, bytes> field (ProtoBytes).
	//
	// An alternative is using protobuf library's any.Any type, which requires individual
	// serialization as well, plus providing a uniquely identifying URL for each message,
	// making it unnecessarily complicated in comparison, if we only want a deterministic
	// encoding of the input.
	bytesByID := map[string][]byte{}
	for id, proto := range protosByID {
		bytes, err := MarshalDeterministic(proto)
		if err != nil {
			return nil, err
		}
		bytesByID[id] = bytes
	}

	return MarshalDeterministic(&protos.ProtosByID{BytesById: bytesByID})
}

// MarshalDeterministic encodes protobuf while enforcing deterministic serialization.
// NOTE: deterministic != canonical, so do not expect this encoding to be
// equal across languages or even versions of golang/protobuf/proto.
// For further reading, see below.
// 	- https://developers.google.com/protocol-buffers/docs/encoding#implications
//	- https://gist.github.com/kchristidis/39c8b310fd9da43d515c4394c3cd9510
func MarshalDeterministic(pb proto.Message) ([]byte, error) {
	buf := &proto.Buffer{}
	buf.SetDeterministic(true)

	err := buf.Marshal(pb)
	return buf.Bytes(), err
}
